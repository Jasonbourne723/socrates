package request

type CreateResource struct {
	Name              string         `json:"name" form:"name" binding:"required"`
	Code              string         `json:"code"  form:"code" binding:"required"`
	Description       string         `json:"description" form:"description"`
	PermissionSpaceId int64          `json:"permission_space_id" form:"permission_space_id" binding:"required"`
	Category          int8           `json:"category" form:"category" binding:"required"`
	Items             []ResourceItem `json:"items" form:"items"`
	Actions           []string       `json:"actions" form:"actions"`
}

type ResourceItem struct {
	Name        string         `json:"name" form:"name"`
	Code        string         `json:"code" form:"code"`
	Value       string         `json:"value" form:"value"`
	Description string         `json:"description" form:"description"`
	Items       []ResourceItem `json:"items" form:"items"`
}

type UpdateResource struct {
	Id                int64          `json:"id" form:"id"  binding:"required"`
	Name              string         `json:"name" form:"name" binding:"required"`
	Code              string         `json:"code"  form:"code" binding:"required"`
	Description       string         `json:"description" form:"description"`
	PermissionSpaceId int64          `json:"permission_space_id" form:"permission_space_id" binding:"required"`
	Items             []ResourceItem `json:"items" form:"items"`
	Actions           []string       `json:"actions" form:"actions"`
}

func (c CreateResource) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":                "名称不能为空",
		"code.required":                "编号不能为空",
		"permission_space_id.required": "权限空间不能为空",
		"category.required":            "类型不能为空",
	}
}

func (c UpdateResource) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"id.required":                  "ID不能为空",
		"name.required":                "名称不能为空",
		"code.required":                "编号不能为空",
		"permission_space_id.required": "权限空间不能为空",
		"category.required":            "类型不能为空",
	}
}
