package tables

import "time"

type EmployeePhone struct {
	EmployeeId       string     `gorm:"primaryKey;type:varchar(36);column:EmployeeId"`
	PhoneId          string     `gorm:"primaryKey;type:varchar(36);column:PhoneId"`
	Note             *string    `gorm:"type:varchar(100);null;column:Note"`
	PhoneTypeId      int64      `gorm:"type:int;not null;column:PhoneTypeId"`
	CreationDate     time.Time  `gorm:"type:datetime;not null;column:CreationDate"`
	CreationEmployee string     `gorm:"type:varchar(36);not null;column:CreationEmployee"`
	ModDate          *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee      *string    `gorm:"type:varchar(36);null;column:ModEmployee"`
}

func (EmployeePhone) TableName() string {
	return "EmployeesPhones"
}
