package request

type Page struct {
	PageIndex int32 `form:"page_index" json:"page_index"`
	PageSize  int32 `form:"page_size" json:"page_size"`
}
