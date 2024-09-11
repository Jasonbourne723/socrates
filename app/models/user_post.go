package models

type UserPost struct {
	ID
	UserId int64 `json:"user_id"`
	PostId int64 `json:"post_id"`
	Timestamps
	SoftDeletes
}
