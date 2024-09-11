package models

type Policy struct {
	ID
	Name        string `json:"name" gorm:"size:20;not null;comment:名称"`
	Description string `json:"description" gorm:"size:100;not null;comment:描述"`
	Timestamps
	SoftDeletes
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
