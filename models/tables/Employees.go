package tables

import "time"

// Table of the database
type Employee struct {
	Id                 string     `gorm:"primaryKey;type:varchar(36);not null;column:Id"`
	Name               string     `gorm:"type:varchar(50);not null;column:Name"`
	LastName           string     `gorm:"type:varchar(100);not null;column:LastName"`
	DocumentId         string     `gorm:"type:varchar(20);not null;column:DocumentId"`
	Email              string     `gorm:"type:varchar(50);not null;column:Email"`
	RoleId             int32      `gorm:"type:int;not null;column:RoleId"`
	CreationEmployeeId string     `gorm:"type:varchar(36);not null;column:CreationEmployee"`
	ModDate            *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployeeId      *string    `gorm:"type:char(36);null;column:ModEmployee"`
}

func (Employee) TableName() string {
	return "Employees"
}
