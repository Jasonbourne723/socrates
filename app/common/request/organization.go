package request

type CreateOrganization struct {
	Name     string `form:"name"  json:"name" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
	ParentId int64  `form:"parent_id" json:"parent_id"`
}

func (p *CreateOrganization) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required": "名称不能为空",
	}
}

type UpdateOrganization struct {
	Id       int64  `form:"id" json:"id" binding:"required"`
	Name     string `form:"name"  json:"name" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
	ParentId int64  `form:"parent_id" json:"parent_id"`
}

func (p *UpdateOrganization) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required": "名称不能为空",
		"id.required":   "id不能为空",
	}
}
