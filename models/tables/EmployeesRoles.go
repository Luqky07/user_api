package tables

import "time"

// Table of the database
type EmployeeRole struct {
	Id               int32      `gorm:"primaryKey;type:int;autoincrement;column:Id"`
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
