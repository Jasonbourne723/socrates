package models

type Organization struct{
	ID 
	Name string `json:"name" gorm:"size:30;not null;comment:名称"`
	ParentId int64 `json:"parent_id"`
	Timestamps
    SoftDeletes
}