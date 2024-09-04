package tables

import "time"

type EmployeeRole struct {
	RoleID           string     `gorm:"primaryKey;type:char(36);column:RoleId"`
	Description      string     `gorm:"type:varchar(50);not null;column:Description"`
	Note             *string    `gorm:"type:varchar(200);null;column:Note"`
	CreationDate     time.Time  `gorm:"type:datetime;not null;column:CreationDate"`
	CreationEmployee string     `gorm:"type:char(36);not null;column:CreationEmployee"`
	ModDate          *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee      *string    `gorm:"type:char(36);null;column:ModEmployee"`
}

func (EmployeeRole) TableName() string {
	return "EmployeesRoles"
}
