package request

type CreatePost struct {
	Code            string  `form:"code" json:"code" binding:"required"`
	Name            string  `form:"name" json:"name" binding:"required"`
	OrganizationIds []int64 `form:"organization_ids" json:"organization_ids"`
}

func (c CreatePost) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required": "名称不能为空",
		"code.required": "编号不能为空",
	}
}

type UpdatePost struct {
	Id int64 `form:"id" json:"id" binding:"required"`
	CreatePost
}

func (c UpdatePost) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"id.required":   "id不能为空",
		"name.required": "名称不能为空",
		"code.required": "编号不能为空",
	}
}
