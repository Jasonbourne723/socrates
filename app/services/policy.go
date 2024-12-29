package services

import (
	"math"

	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/models"
	"github.com/Jasonbourne723/socrates/global"
)

type policyService struct {
}

type IPolicyService interface {
	PageList(pageIndex int32, pageSize int32) (pages response.Page[response.Policy], err error)
	List() (list []*response.Policy, err error)
	Create(req *request.CreatePolicy) (res *response.Policy, err error)
	Update(req *request.UpdatePolicy) (res *response.Policy, err error)
	Delete(id int64) (err error)
}

func NewPolicyService() IPolicyService {
	return &policyService{}
}

func (p *policyService) PageList(pageIndex int32, pageSize int32) (pages response.Page[response.Policy], err error) {
	var policys []models.Policy
	if err = global.App.DB.Offset((int(pageIndex) - 1) * int(pageSize)).Limit(int(pageSize)).Find(&policys).Error; err != nil {
		return
	}

	rows := []response.Policy{}
	for _, item := range policys {
		rows = append(rows, *MapToPolicyResponse(&item))
	}

	var count int64
	err = global.App.DB.Find(&models.Role{}).Count(&count).Error

	pages = response.Page[response.Policy]{
		Rows:       rows,
		PageIndex:  int32(pageIndex),
		PageSize:   int32(pageSize),
		TotalCount: count,
		TotalPage:  int64(math.Ceil(float64(count) / float64(pageSize))),
	}

	return
}
func (p *policyService) List() (list []*response.Policy, err error) {
	var policys []models.Policy
	if err = global.App.DB.Find(&policys).Error; err != nil {
		return
	}

	list = []*response.Policy{}
	for _, item := range policys {
		list = append(list, MapToPolicyResponse(&item))
	}
	return
}

func (p *policyService) Create(req *request.CreatePolicy) (res *response.Policy, err error) {
	exists := models.Policy{}
	result := global.App.DB.Where("name = ?", req.Name).First(&exists)
	if result.RowsAffected != 0 {
		return nil, global.Errors.NameDuplicateError
	}
	entity := models.Policy{Name: req.Name}
	if err = global.App.DB.Create(&entity).Error; err != nil {
		return
	}

	for _, resourceDto := range req.Resources {
		resource := models.PolicyResource{
			PermissionSpaceId: resourceDto.PermissionSpaceId,
			ResourceId:        resourceDto.ResourceId,
			Effect:            resourceDto.Effect,
			PolicyId:          entity.Id,
		}
		if err = global.App.DB.Create(&resource).Error; err != nil {
			return
		}

		resourceItems := []models.PolicyResourceItem{}
		for _, item := range resourceDto.Items {
			resourceItems = append(resourceItems, models.PolicyResourceItem{
				PolicyId:            entity.Id,
				PolicyResourceId:    resource.Id,
				ResourceItemId:      item.ResourceItemId,
				ResourceItemActions: item.ResourceItemActions,
			})
		}
		if err = global.App.DB.Create(&resourceItems).Error; err != nil {
			return
		}
	}

	return
}

func (p *policyService) Update(req *request.UpdatePolicy) (res *response.Policy, err error) {
	var exists models.Policy
	result := global.App.DB.First(&exists, req.Id)
	if result.Error != nil {
		return
	}
	if result.RowsAffected == 0 {
		err = global.Errors.RecordNotFoundError
		return
	}

	exists.Name = req.Name

	err = global.App.DB.Save(&exists).Error
	if err != nil {
		return
	}
	return
}

func (p *policyService) Delete(id int64) (err error) {
	err = global.App.DB.Delete(&models.Policy{}, id).Error
	return err
}

func MapToPolicyResponse(m *models.Policy) *response.Policy {
	return &response.Policy{
		Id:          m.Id,
		Name:        m.Name,
		Description: m.Description,
	}
}
