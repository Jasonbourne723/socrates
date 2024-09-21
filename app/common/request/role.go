package request

type CreateRole struct {
	Code              string `form:"code" json:"code" binding:"required"`
	Name              string `form:"name" json:"name" binding:"required"`
	PermissionSpaceId int64  `form:"permission_space_id" json:"permission_space_id" binding:"required"`
}

func (dto CreateRole) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required": "名称不能为空",
		"code.required": "编号不能为空",
	}
}

type UpdateRole struct {
	Id                int64  `form:"id" json:"id" binding:"required"`
	Code              string `form:"code" json:"code" binding:"required"`
	Name              string `form:"name" json:"name" binding:"required"`
	PermissionSpaceId int64  `form:"permission_space_id" json:"permission_space_id" binding:"required"`
}

func (dto UpdateRole) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":       "名称不能为空",
		"code.required":       "编号不能为空",
		"permission_space_id": "权限空间不能为空",
	}
}
