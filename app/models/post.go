package models

type Post struct {
	ID
	Code string `json:"code" gorm:"size:20;not null;comment:代码"`
	Name string `json:"name" gorm:"size:20;not null;comment:名称"`
	Timestamps
	SoftDeletes
}
