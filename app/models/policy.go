package models

type Policy struct {
	ID
	Name        string `json:"name" gorm:"size:20;not null;comment:名称"`
	Description string `json:"description" gorm:"size:100;not null;comment:描述"`
	Timestamps
	SoftDeletes
}

type PolicyAuth struct {
	ID
	PolicyId   int64 `json:"policy_id"`
	TargetType int8  `json:"target_type" gorm:"comment:0用户1角色2部门"`
	TargetId   int64 `json:"target_id"`
}

type PolicyResource struct {
	ID
	PolicyId          int64 `json:"policy_id"`
	PermissionSpaceId int64 `json:"permission_space_id"`
	ResourceId        int64 `json:"resource_id"`
	Effect            int8  `json:"effect" gorm:"comment:效果0/1"`
	Timestamps
	SoftDeletes
}

type PolicyResourceItem struct {
	ID
	PolicyId            int64  `json:"policy_id"`
	PolicyResourceId    int64  `json:"policy_resouce_id"`
	ResourceItemId      int64  `json:"resource_item_id"`
	ResourceItemActions string `json:"resource_item_actions" gorm:"size:255;not null"`
}
