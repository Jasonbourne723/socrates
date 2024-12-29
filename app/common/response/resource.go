package response

type Resource struct {
	Id                int64          `json:"id"`
	Name              string         `json:"name" form:"name" binding:"required"`
	Code              string         `json:"code"  form:"code" binding:"required"`
	Description       string         `json:"description" form:"description"`
	PermissionSpaceId int64          `json:"permission_space_id" form:"permission_space_id" binding:"required"`
	Category          int8           `json:"category" form:"category" binding:"required"`
	Items             []ResourceItem `json:"items" form:"items"`
	Actions           []string       `json:"actions" form:"actions"`
}

type ResourceItem struct {
	Id          int64          `json:"id"`
	Name        string         `json:"name" form:"name"`
	Code        string         `json:"code" form:"code"`
	Value       string         `json:"value" form:"value"`
	Description string         `json:"description" form:"description"`
	Items       []ResourceItem `json:"items" form:"items"`
}
