package models

type UserRole struct {
	ID
	UserId int64 `json:"user_id"`
	RoleId int64 `json:"role_id"`
	Timestamps
    SoftDeletes
}
