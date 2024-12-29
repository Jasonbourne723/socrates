package response

type PermissionSpace struct {
	Id            int64  `form:"id" json:"id" binding:"required"`
	Name          string `form:"name" json:"name" binding:"required"`
	Code          string `form:"code" json:"code" binding:"required"`
	Description   string `form:"description" json:"description" binding:"required"`
	ApplicationId int64  `form:"application_id" json:"application_id" binding:"required"`
}
