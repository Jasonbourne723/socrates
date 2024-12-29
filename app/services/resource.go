package services

import (
	"math"

	"github.com/Jasonbourne723/socrates/app/common/mapster"
	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/models"
	"github.com/Jasonbourne723/socrates/global"
)

func NewResourceService() IResourceSerivce {
	return &ResourceService{}
}

type ResourceService struct {
}

type IResourceSerivce interface {
	Create(params request.CreateResource) (err error)
	Delete(id int64) (err error)
	PageList(pageIndex int32, pageSize int32) (pages response.Page[response.Resource], err error)
	Update(params request.UpdateResource) (err error)
	List() (rows []response.Resource, err error)
	GetOne(resourceId int64) (res response.Resource, err error)
}

func (i *ResourceService) PageList(pageIndex int32, pageSize int32) (pages response.Page[response.Resource], err error) {
	var Resources []models.Resource
	if err = global.App.DB.Offset((int(pageIndex) - 1) * int(pageSize)).Limit(int(pageSize)).Find(&Resources).Error; err != nil {
		return
	}
	rows := []response.Resource{}
	for _, resource := range Resources {
		resourceItems := []models.ResourceItem{}
		if err = global.App.DB.Where("resource_id = ?", resource.Id).Find(&resourceItems).Error; err != nil {
			return
		}
		resourceActions := []models.ResourceAction{}
		if err = global.App.DB.Where("resource_id = ?", resource.Id).Find(&resourceActions).Error; err != nil {
			return
		}
		rows = append(rows, *MapToResourceResponse(&resource, resourceItems, resourceActions))
	}

	var count int64
	err = global.App.DB.Find(&models.Resource{}).Count(&count).Error

	pages = response.Page[response.Resource]{
		Rows:       rows,
		PageIndex:  pageIndex,
		PageSize:   pageSize,
		TotalCount: count,
		TotalPage:  int64(math.Ceil(float64(count) / float64(pageSize))),
	}
	return
}

func (i *ResourceService) GetOne(resourceId int64) (res response.Resource, err error) {
	var resource models.Resource
	if err = global.App.DB.First(&resource, resourceId).Error; err != nil {
		return
	}

	resourceItems := []models.ResourceItem{}
	if err = global.App.DB.Where("resource_id = ?", resource.Id).Find(&resourceItems).Error; err != nil {
		return
	}
	resourceActions := []models.ResourceAction{}
	if err = global.App.DB.Where("resource_id = ?", resource.Id).Find(&resourceActions).Error; err != nil {
		return
	}
	res = *MapToResourceResponse(&resource, resourceItems, resourceActions)
	return
}

func (i *ResourceService) List() (rows []response.Resource, err error) {
	var Resources []models.Resource
	if err = global.App.DB.Find(&Resources).Error; err != nil {
		return
	}
	rows = []response.Resource{}
	for _, resource := range Resources {
		resourceItems := []models.ResourceItem{}
		if err = global.App.DB.Where("resource_id = ?", resource.Id).Find(&resourceItems).Error; err != nil {
			return
		}
		resourceActions := []models.ResourceAction{}
		if err = global.App.DB.Where("resource_id = ?", resource.Id).Find(&resourceActions).Error; err != nil {
			return
		}
		rows = append(rows, *MapToResourceResponse(&resource, resourceItems, resourceActions))
	}
	return
}

func saveResourceItemsForTree(resourceId int64, parentId int64, items []request.ResourceItem) error {
	for _, item := range items {
		// 插入当前资源明细
		resourceItem := models.ResourceItem{
			Name:        item.Name,
			Code:        item.Code,
			Value:       item.Value,
			Description: item.Description,
			ParentId:    parentId,   // 父级ID
			ResourceId:  resourceId, // 资源ID
		}

		if err := global.App.DB.Create(&resourceItem).Error; err != nil {
			return err
		}

		// 如果当前明细有子明细，递归插入
		if len(item.Items) > 0 {
			if err := saveResourceItemsForTree(resourceId, resourceItem.Id, item.Items); err != nil {
				return err
			}
		}
	}
	return nil
}

func (i *ResourceService) Create(params request.CreateResource) (err error) {

	existsResource := models.Resource{}
	result := global.App.DB.Where("code = ? ", params.Code).Or("name = ?", params.Name).First(&existsResource)
	if result.RowsAffected != 0 {
		if existsResource.Code == params.Code {
			return global.Errors.CodeDuplicateError
		} else {
			return global.Errors.NameDuplicateError
		}
	}

	entity := models.Resource{
		Name:              params.Name,
		Code:              params.Code,
		PermissionSpaceId: params.PermissionSpaceId,
		Description:       params.Description,
		Category:          params.Category,
	}
	//资源
	if err = global.App.DB.Create(&entity).Error; err != nil {
		return
	}
	//资源明细
	if params.Category == 1 {
		if err = saveResourceItemsForTree(entity.Id, 0, params.Items); err != nil {
			return
		}
	} else if params.Category == 2 {
		if err = saveResourceItemsForArray(entity.Id, params.Items); err != nil {
			return
		}
	}
	//资源操作
	if err = saveResourceActions(entity.Id, params.Actions); err != nil {
		return
	}

	return
}

func (i *ResourceService) Delete(id int64) (err error) {
	if err = global.App.DB.Delete(&models.Resource{}, id).Error; err != nil {
		return
	}
	if err = global.App.DB.Where("resource_id = ?", id).Delete(&models.ResourceItem{}).Error; err != nil {
		return
	}
	if err = global.App.DB.Where("resource_id = ?", id).Delete(&models.ResourceAction{}).Error; err != nil {
		return
	}
	return err
}

func (i *ResourceService) Update(params request.UpdateResource) (err error) {
	var existResource models.Resource
	result := global.App.DB.First(&existResource, params.Id)
	if result.Error != nil {
		return
	}
	if result.RowsAffected == 0 {
		err = global.Errors.RecordNotFoundError
		return
	}

	existResource.Code = params.Code
	existResource.Name = params.Name
	existResource.PermissionSpaceId = params.PermissionSpaceId
	existResource.Description = params.Description

	err = global.App.DB.Save(&existResource).Error
	if err != nil {
		return
	}
	//资源明细
	if err = global.App.DB.Where("resource_id = ?", existResource.Id).Delete(&models.ResourceItem{}).Error; err != nil {
		return
	}
	if existResource.Category == 1 {
		if err = saveResourceItemsForTree(existResource.Id, 0, params.Items); err != nil {
			return
		}
	} else if existResource.Category == 2 {
		if err = saveResourceItemsForArray(existResource.Id, params.Items); err != nil {
			return
		}
	}
	//资源操作
	if err = global.App.DB.Where("resource_id = ?", existResource.Id).Delete(&models.ResourceAction{}).Error; err != nil {
		return
	}
	if err = saveResourceActions(existResource.Id, params.Actions); err != nil {
		return
	}
	return
}

func buildResourceItemTree(items []models.ResourceItem, parentId int64) []response.ResourceItem {
	var result []response.ResourceItem
	for _, item := range items {
		if item.ParentId == parentId {
			// 找到当前项的子项，递归构建树
			nestedItem := response.ResourceItem{
				Id:          item.Id,
				Name:        item.Name,
				Code:        item.Code,
				Value:       item.Value,
				Description: item.Description,
				Items:       buildResourceItemTree(items, item.Id), // 递归查找子项
			}
			result = append(result, nestedItem)
		}
	}
	return result
}

func MapToResourceResponse(resource *models.Resource, resourceItems []models.ResourceItem, actions []models.ResourceAction) *response.Resource {

	items := []response.ResourceItem{}
	if resource.Category == 1 {
		items = buildResourceItemTree(resourceItems, 0)

	} else if resource.Category == 2 {
		items = mapster.Map(resourceItems, func(resouceItem models.ResourceItem) response.ResourceItem {
			return response.ResourceItem{
				Id:          resouceItem.Id,
				Name:        resouceItem.Name,
				Code:        resouceItem.Code,
				Value:       resouceItem.Value,
				Description: resouceItem.Description,
			}
		})
	}

	return &response.Resource{
		Id:                resource.Id,
		Name:              resource.Name,
		Code:              resource.Code,
		PermissionSpaceId: resource.PermissionSpaceId,
		Description:       resource.Description,
		Category:          resource.Category,
		Actions: mapster.Map(actions, func(a models.ResourceAction) string {
			return a.Name
		}),
		Items: items,
	}
}
func saveResourceItemsForArray(resourceId int64, items []request.ResourceItem) error {
	resourceitems := []models.ResourceItem{}
	for _, item := range items {
		resourceitems = append(resourceitems, models.ResourceItem{
			Name:        item.Name,
			Code:        item.Code,
			Value:       item.Value,
			ResourceId:  resourceId,
			Description: item.Description,
			ParentId:    0,
		})
	}
	return global.App.DB.Create(&resourceitems).Error
}

func saveResourceActions(resourceId int64, actions []string) error {
	resourceActions := []models.ResourceAction{}
	for _, action := range actions {
		resourceActions = append(resourceActions, models.ResourceAction{
			ResourceId: resourceId,
			Name:       action,
		})
	}
	return global.App.DB.Create(&resourceActions).Error
}
