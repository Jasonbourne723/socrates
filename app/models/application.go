package models

type Application struct {
	ID
	Name        string `json:"name" gorm:"size:20;not null;comment:名称"`
	Description string `json:"description" gorm:"size:100;not null;comment:描述"`
	AppKey      string `json:"app_key" gorm:"size:50;not null;comment:appkey"`
	AppSecret   string `json:"app_secret" gorm:"size:100;not null;comment:密钥"`
	CallbackUrl string `json:"callback_url" gorm:"size:255;not null;comment:登录回调地址"`
	Timestamps
	SoftDeletes
}
