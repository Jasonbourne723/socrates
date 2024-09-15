package response

type Post struct {
	Id              int64   `form:"id" json:"id"`
	Code            string  `form:"code" json:"code"`
	Name            string  `form:"name" json:"name"`
	OrganizationIds []int64 `form:"organization_ids" json:"organization_ids"`
}
