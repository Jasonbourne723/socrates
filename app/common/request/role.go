package request

type CreaeteRole struct {
	Code              string `form:"code" json:"code" binding:"required"`
	Name              string `form:"name" json:"name" binding:"required"`
	PermissionSpaceId int64  `form:permission_space_id json:"permission_space_id" binding:"required"`
}

func (register CreaeteRole) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required": "名称不能为空",
		"code.required": "编号不能为空",
	}
}
