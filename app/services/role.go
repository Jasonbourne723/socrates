package services

import (
	"math"

	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/models"
	"github.com/Jasonbourne723/socrates/global"
)

func NewRoleService() IRoleSerivce {
	return &RoleService{}
}

type RoleService struct {
}

type IRoleSerivce interface {
	Create(params request.CreaeteRole) (role *response.Role, err error)
	Delete(id int64) (err error)
	PageList(pageIndex int32, pageSize int32) (pages response.Page[response.Role], err error)
}

func (i *RoleService) Create(params request.CreaeteRole) (role *response.Role, err error) {

	existsRole := models.Role{}
	result := global.App.DB.Where("code = ? ", params.Code).Or("name = ?", params.Name).First(&existsRole)
	if result.RowsAffected != 0 {
		if existsRole.Code == params.Code {
			return nil, global.Errors.CodeDuplicateError
		} else {
			return nil, global.Errors.NameDuplicateError
		}
	}

	entity := models.Role{Name: params.Name, Code: params.Code, PermissionSpaceId: params.PermissionSpaceId}
	err = global.App.DB.Create(&entity).Error
	role = MapToRoleResponse(&entity)
	return
}

func (i *RoleService) Delete(id int64) (err error) {
	err = global.App.DB.Delete(&models.Role{}, id).Error
	return err
}

func (i *RoleService) PageList(pageIndex int32, pageSize int32) (pages response.Page[response.Role], err error) {
	var roles []models.Role
	if err = global.App.DB.Offset((int(pageIndex) - 1) * int(pageSize)).Limit(int(pageSize)).Find(&roles).Error; err != nil {
		return
	}
	rows := []response.Role{}
	for _, item := range roles {
		rows = append(rows, *MapToRoleResponse(&item))
	}

	var count int64
	err = global.App.DB.Find(&models.Role{}).Count(&count).Error

	pages = response.Page[response.Role]{
		Rows:       rows,
		PageIndex:  pageIndex,
		PageSize:   pageSize,
		TotalCount: count,
		TotalPage:  int64(math.Ceil(float64(count) / float64(pageSize))),
	}
	return
}

func MapToRoleResponse(entity *models.Role) *response.Role {
	return &response.Role{
		Id:                entity.Id,
		Name:              entity.Name,
		Code:              entity.Code,
		PermissionSpaceId: entity.PermissionSpaceId,
	}
}
