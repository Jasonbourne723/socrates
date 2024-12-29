package response

import "time"

type User struct {
	Id             int64     `json:"id" form:"id"`
	Name           string    `json:"name" form:"name"`
	Mobile         string    `json:"mobile" form:"mobile"`
	RoleIds        []int64   `json:"role_ids"`
	OrganizationId int64     `json:"organization_id"`
	PostId         int64     `json:"post_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
