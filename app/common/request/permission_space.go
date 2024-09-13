package request

type CreatePermissionSpace struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Code        string `form:"code" json:"code" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
}

func (dto CreatePermissionSpace) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":        "名称不能为空",
		"code.required":        "编号不能为空",
		"description.required": "描述不能为空",
	}
}

type UpdatePermissionSpace struct {
	Id int64 `form:"id" json:"id" binding:"required" `
	CreatePermissionSpace
}

func (dto UpdatePermissionSpace) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":        "名称不能为空",
		"code.required":        "编号不能为空",
		"description.required": "描述不能为空",
		"id.required":          "ID不能为空",
	}
}
