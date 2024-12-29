package models

type PermissionSpace struct {
	ID
	Name        string `json:"name" gorm:"size:20;not null;comment:名称"`
	Code        string `json:"code" gorm:"size:50;not null;comment:编号"`
	Description string `json:"description" gorm:"size:100;not null;comment:描述"`
	Timestamps
	SoftDeletes
}


