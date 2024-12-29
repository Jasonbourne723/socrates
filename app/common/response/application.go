package response

type Application struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AppKey      string `json:"app_key"`
	AppSecret   string `json:"app_secret"`
	CallbackUrl string `json:"callback_url"`
}

type SsoApplication struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AppKey      string `json:"app_key"`
	AppSecret   string `json:"app_secret"`
	CallbackUrl string `json:"callback_url"`
}
