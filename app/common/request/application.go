package request

type CreateApplication struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description"`
	CallbackUrl string `form:"description" json:"callback_url" `
}

func (dto CreateApplication) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required": "名称不能为空",
	}
}

type UpdateApplication struct {
	Id int64 `form:"id" json:"id"`
	CreateApplication
}

func (dto UpdateApplication) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required": "名称不能为空",
		"id.required":   "id不能为空",
	}
}
