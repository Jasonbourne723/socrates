package services

import (
	"math"

	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/models"
	"github.com/Jasonbourne723/socrates/global"
)

type PermissionSpaceService struct {
}

type IPermissionSpaceService interface {
	PageList(pageIndex int32, pageSize int32) (pages response.Page[response.PermissionSpace], err error)
	List() (list []response.PermissionSpace, err error)
	Create(req *request.CreatePermissionSpace) (res *response.PermissionSpace, err error)
	Update(req *request.UpdatePermissionSpace) (res *response.PermissionSpace, err error)
	Delete(id int64) (err error)
}

func NewPermissionSpaceService() IPermissionSpaceService {
	return &PermissionSpaceService{}
}

func (p *PermissionSpaceService) PageList(pageIndex int32, pageSize int32) (pages response.Page[response.PermissionSpace], err error) {

	var permissionSpaces []models.PermissionSpace
	if err = global.App.DB.Offset((int(pageIndex) - 1) * int(pageSize)).Limit(int(pageSize)).Find(&permissionSpaces).Error; err != nil {
		return
	}

	rows := []response.PermissionSpace{}
	for _, item := range permissionSpaces {
		rows = append(rows, *MapToPermissionSpaceResponse(&item))
	}

	var count int64
	err = global.App.DB.Find(&models.Role{}).Count(&count).Error

	pages = response.Page[response.PermissionSpace]{
		Rows:       rows,
		PageIndex:  int32(pageIndex),
		PageSize:   int32(pageSize),
		TotalCount: count,
		TotalPage:  int64(math.Ceil(float64(count) / float64(pageSize))),
	}

	return
}

func (p *PermissionSpaceService) List() (list []response.PermissionSpace, err error) {
	var permissionSpaces []models.PermissionSpace
	if err = global.App.DB.Find(&permissionSpaces).Error; err != nil {
		return
	}
	list = []response.PermissionSpace{}
	for _, item := range permissionSpaces {
		list = append(list, *MapToPermissionSpaceResponse(&item))
	}
	return
}

func (p *PermissionSpaceService) Create(req *request.CreatePermissionSpace) (res *response.PermissionSpace, err error) {
	exists := models.PermissionSpace{}
	result := global.App.DB.Where("code = ? ", req.Code).Or("name = ?", req.Name).First(&exists)
	if result.RowsAffected != 0 {
		if exists.Code == req.Code {
			return nil, global.Errors.CodeDuplicateError
		} else {
			return nil, global.Errors.NameDuplicateError
		}
	}

	entity := models.PermissionSpace{Name: req.Name, Code: req.Code, Description: req.Description}
	err = global.App.DB.Create(&entity).Error
	res = MapToPermissionSpaceResponse(&entity)
	return
}

func (p *PermissionSpaceService) Update(req *request.UpdatePermissionSpace) (res *response.PermissionSpace, err error) {
	var exists models.PermissionSpace
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
	exists.Description = req.Description

	err = global.App.DB.Save(&exists).Error
	if err != nil {
		return
	}
	res = MapToPermissionSpaceResponse(&exists)
	return
}

func (p *PermissionSpaceService) Delete(id int64) (err error) {
	err = global.App.DB.Delete(&models.PermissionSpace{}, id).Error
	return err
}

func MapToPermissionSpaceResponse(m *models.PermissionSpace) *response.PermissionSpace {
	return &response.PermissionSpace{
		Id:          m.Id,
		Name:        m.Name,
		Code:        m.Code,
		Description: m.Description,
	}
}
