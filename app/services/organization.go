package services

import (
	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/models"
	"github.com/Jasonbourne723/socrates/global"
)

type OrganizationService struct {
}

type IOrganizationService interface {
	List() (organizationDtos []*response.Organization, err error)
	All() (list []*response.OrganizationNode, err error)
	Create(req *request.CreateOrganization) (res *response.Organization, err error)
	Update(req *request.UpdateOrganization) (res *response.Organization, err error)
	Delete(id int64) (err error)
}

func NewOrganizationService() IOrganizationService {
	return &OrganizationService{}
}

func (p *OrganizationService) All() (roots []*response.OrganizationNode, err error) {
	var organizations []models.Organization
	if err = global.App.DB.Find(&organizations).Error; err != nil {
		return
	}
	idMap := make(map[int64]*response.OrganizationNode)
	roots = []*response.OrganizationNode{}
	for i := range organizations {
		entity := MapToOrganizationNodeResponse(&organizations[i])
		idMap[entity.Id] = entity
	}

	for i := range organizations {
		org := &organizations[i]
		if org.ParentId == 0 {
			// 如果是根节点
			roots = append(roots, idMap[org.Id])
		} else {
			// 如果有父节点，找到父节点并将其添加到父节点的children中
			parent, exists := idMap[org.ParentId]
			if exists {
				parent.Items = append(parent.Items, idMap[org.Id])
			}
		}
	}

	return roots, err
}

func (p *OrganizationService) List() (organizationDtos []*response.Organization, err error) {
	var organizations []models.Organization
	if err = global.App.DB.Find(&organizations).Error; err != nil {
		return
	}
	organizationDtos = []*response.Organization{}
	for _, ea := range organizations {
		organizationDtos = append(organizationDtos, MapToOrganizationResponse(&ea))
	}
	return
}

func (p *OrganizationService) Create(req *request.CreateOrganization) (res *response.Organization, err error) {
	exists := models.Organization{}
	result := global.App.DB.Where("code = ? ", req.Code).Or("name = ?", req.Name).First(&exists)
	if result.RowsAffected != 0 {
		if exists.Code == req.Code {
			return nil, global.Errors.CodeDuplicateError
		} else {
			return nil, global.Errors.NameDuplicateError
		}
	}

	entity := models.Organization{Name: req.Name, Code: req.Code, ParentId: req.ParentId}
	err = global.App.DB.Create(&entity).Error
	res = MapToOrganizationResponse(&entity)
	return
}

func (p *OrganizationService) Update(req *request.UpdateOrganization) (res *response.Organization, err error) {
	var exists models.Organization
	result := global.App.DB.First(&exists, req.Id)
	if result.Error != nil {
		return
	}
	if result.RowsAffected == 0 {
		err = global.Errors.RecordNotFoundError
		return
	}

	exists.Code = req.Code
	exists.Name = req.Name
	exists.ParentId = req.ParentId

	err = global.App.DB.Save(&exists).Error
	if err != nil {
		return
	}
	res = MapToOrganizationResponse(&exists)
	return
}

func (p *OrganizationService) Delete(id int64) (err error) {
	err = global.App.DB.Delete(&models.Organization{}, id).Error
	return err
}

func MapToOrganizationResponse(m *models.Organization) *response.Organization {
	return &response.Organization{
		Id:       m.Id,
		Name:     m.Name,
		Code:     m.Code,
		ParentId: m.ParentId,
	}
}

func MapToOrganizationNodeResponse(m *models.Organization) *response.OrganizationNode {
	return &response.OrganizationNode{
		Id:       m.Id,
		Name:     m.Name,
		Code:     m.Code,
		ParentId: m.ParentId,
		Items:    []*response.OrganizationNode{},
	}
}
