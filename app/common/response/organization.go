package response

type OrganizationNode struct {
	Id    int64              `json:"id" form:"id"`
	Code  string             `form:"code" json:"code" binding:"required"`
	Name  string             `json:"name" form:"name" `
	Items []OrganizationNode `json:"items" form:"items"`
}

type Organization struct {
	Id       int64  `json:"id" form:"id"`
	Code     string `form:"code" json:"code" binding:"required"`
	Name     string `json:"name" form:"name" `
	ParentId int64  `json:"parent_id" form:"parent_id"`
}
