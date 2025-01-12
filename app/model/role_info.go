package model

type RoleInfo struct {
	BaseModel
	Name            string `gorm:"not null"`
	Description     string
	PermissionInfos []PermissionInfo `gorm:"many2many:roleInfo_permissionInfo;"`
	Users           []User
}
