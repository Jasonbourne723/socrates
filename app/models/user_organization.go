package models

type UserOrganization struct {
	ID
	UserId         int64 `json:"user_id"`
	OrganizationId int64 `json:"organization_id"`
	Timestamps
    SoftDeletes
}
