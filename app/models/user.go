package models

import "strconv"

type User struct {
	ID
	Name         string `json:"name" gorm:"size:30;not null;comment:用户名称"`
	Mobile       string `json:"mobile" gorm:"size:24;not null;index;comment:用户手机号"`
	Password     string `json:"-" gorm:"not null;default:'';comment:用户密码"`
	GithubOpenid int64  `json:"github_openid" gorm:"comment:github openid"`
	Avatar       string `json:"avatar" gorm:"size:255;not null;comment:头像"`
	Timestamps
	SoftDeletes
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.Id))
}
