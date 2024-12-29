package request

type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":     "用户名称不能为空",
		"mobile.required":   "手机号码不能为空",
		"mobile.mobile":     "手机号码格式不正确",
		"password.required": "用户密码不能为空",
	}
}

type Login struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"mobile.required":   "手机号码不能为空",
		"mobile.mobile":     "手机号码格式不正确",
		"password.required": "用户密码不能为空",
	}
}

type GithubLogin struct {
	Code string `form:"code" json:"code" binding:"required"`
}

func (login GithubLogin) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"code.required": "授权码不能为空",
	}
}

type CreateUser struct {
	Name           string  `form:"name" json:"name" binding:"required"`
	Mobile         string  `form:"mobile" json:"mobile" binding:"required,mobile"`
	RoleIds        []int64 `json:"role_ids" form:"role_ids"`
	OrganizationId int64   `json:"organization_id" form:"organization_id"`
	PostId         int64   `json:"post_id" form:"post_id"`
}

func (u CreateUser) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":   "用户名称不能为空",
		"mobile.required": "手机号码不能为空",
		"mobile.mobile":   "手机号码格式不正确",
	}
}

type UpdateUser struct {
	Id             int64   `json:"id" form:"id"  binding:"required"`
	Name           string  `form:"name" json:"name" binding:"required"`
	Mobile         string  `form:"mobile" json:"mobile" binding:"required,mobile"`
	RoleIds        []int64 `json:"role_ids" form:"role_ids"`
	OrganizationId int64   `json:"organization_id" form:"organization_id"`
	PostId         int64   `json:"post_id" form:"post_id"`
}

func (u UpdateUser) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":   "用户名称不能为空",
		"mobile.required": "手机号码不能为空",
		"mobile.mobile":   "手机号码格式不正确",
	}
}
