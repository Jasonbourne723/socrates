package models

type PostOrganization struct {
	ID
	OrganizationId int64 `json:"organization_id"`
	PostId         int64 `json:"post_id"`
	Timestamps
    SoftDeletes
}
