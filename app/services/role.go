package services

import (
	"github.com/Jasonbourne723/socrates/app/common/request"
	"github.com/Jasonbourne723/socrates/app/common/response"
	"github.com/Jasonbourne723/socrates/app/models"
	"github.com/Jasonbourne723/socrates/global"
)

type RoleService struct {
}

type IRoleSerivce interface {
	Create(params request.CreaeteRole) (err error)
}

func (i *RoleService) Create(params request.CreaeteRole) (err error, role response.Role) {

	existsRole := models.Role{}
	result := global.App.DB.Where("code = ? ", params.Code).Or("name = ?", params.Name).First(&existsRole)
	if result.RowsAffected != 0 {
		if existsRole.Code == params.Code {
			err = global.Errors.CodeDuplicateError
		} else {
			err = global.Errors.NameDuplicateError
		}
		return
	}

	entity := models.Role{Name: params.Name, Code: params.Code, PermissionSpaceId: params.PermissionSpaceId}
	err = global.App.DB.Create(&entity).Error
	role = MapToRoleResponse(&entity)
	return
}

func MapToRoleResponse(entity *models.Role) response.Role {
	return response.Role{
		Id:                entity.Id,
		Name:              entity.Name,
		Code:              entity.Code,
		PermissionSpaceId: entity.PermissionSpaceId,
	}
}
