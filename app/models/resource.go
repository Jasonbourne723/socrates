package models

type Resource struct {
	ID
	Name              string `json:"name" gorm:"size:20;not null;comment:名称"`
	Code              string `json:"code" gorm:"size:50;not null;comment:编号"`
	Description       string `json:"description" gorm:"size:100;not null;comment:描述"`
	PermissionSpaceId int64  `json:"permission_space_id"`
	Category          int8   `json:"category"`
	Timestamps
	SoftDeletes
}

type ResourceItem struct {
	ID
	Name        string `json:"name" gorm:"size:20;not null;comment:名称"`
	Code        string `json:"code" gorm:"size:50;not null;comment:编号"`
	Description string `json:"description" gorm:"size:100;not null;comment:描述"`
	ResourceId  int64  `json:"resource_id"`
	Timestamps
	SoftDeletes
}

type ResourceAction struct {
	ID
	Name        string `json:"name" gorm:"size:20;not null;comment:名称"`
	Description string `json:"description" gorm:"size:100;not null;comment:描述"`
	ResourceId  int64  `json:"resource_id"`
}
