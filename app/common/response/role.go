package response

type Role struct {
	Id                int64
	Code              string `json:"code" gorm:"size:30;not null;comment:编号"`
	Name              string `json:"name" gorm:"size:30;not null;comment:用户名称"`
	PermissionSpaceId int64  `json:"permission_space_id"`
}
