package models

type Resource struct {
	ID
	Name              string `json:"name" gorm:"size:20;not null;comment:名称"`
	Code              string `json:"code" gorm:"size:50;not null;comment:编号"`
	Description       string `json:"description" gorm:"size:100;not null;comment:描述"`
	PermissionSpaceId int64  `json:"permission_space_id"`
	Category          int8   `json:"category" gorm:"not null;comment:1树形2数组"`
	Timestamps
	SoftDeletes
}

type ResourceItem struct {
	ID
	Name        string `json:"name" gorm:"size:20;not null;comment:名称"`
	Code        string `json:"code" gorm:"size:50;not null;comment:编号"`
	Value       string `json:"value" gorm:"size:50;not null:comment:值"`
	Description string `json:"description" gorm:"size:50;not null;comment:描述"`
	ParentId    int64  `json:"parent_id"`
	ResourceId  int64  `json:"resource_id"`
	Timestamps
	SoftDeletes
}

type ResourceAction struct {
	ID
	Name       string `json:"name" gorm:"size:20;not null;comment:名称"`
	ResourceId int64  `json:"resource_id"`
}
