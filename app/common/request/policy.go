package request

type CreatePolicy struct {
	Name        string           `json:"name" form:"name"`
	Description string           `json:"description" form:"description"`
	Resources   []PolicyResource `json:"resources" form:"resources"`
}

type UpdatePolicy struct {
	Id          int64            `json:"id" form:"id"`
	Name        string           `json:"name" form:"name"`
	Description string           `json:"description" form:"description"`
	Resources   []PolicyResource `json:"resources" form:"resources"`
}

type PolicyResource struct {
	PermissionSpaceId int64                `json:"permission_space_id" form:"permission_space_id"`
	ResourceId        int64                `json:"resource_id" form:"resource_id"`
	Effect            int8                 `json:"effect" form:"effect"`
	Items             []PolicyResourceItem `json:"items" form:"items"`
}

type PolicyResourceItem struct {
	ResourceItemId      int64  `json:"resource_item_id" form:"resource_item_id"`
	ResourceItemName    string `json:"resource_item_name" form:"resource_item_name"`
	ResourceItemActions string `json:"resource_item_actions" form:"resource_item_actions"`
}

func (c CreatePolicy) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required": "名称不能为空",
	}
}
func (c UpdatePolicy) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required": "名称不能为空",
	}
}
