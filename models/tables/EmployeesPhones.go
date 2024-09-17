package tables

import (
	"time"

	"gorm.io/gorm"
)

// Table of the database
type EmployeePhone struct {
	EmployeeId       string     `gorm:"primaryKey;type:char(36);column:EmployeeId"`
	PhoneId          string     `gorm:"primaryKey;type:char(36);column:PhoneId"`
	Note             *string    `gorm:"type:varchar(100);null;column:Note"`
	PhoneTypeId      int32      `gorm:"type:int;not null;column:PhoneTypeId"`
	CreationDate     time.Time  `gorm:"type:datetime;not null;column:CreationDate"`
	CreationEmployee string     `gorm:"type:char(36);not null;column:CreationEmployee"`
	ModDate          *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee      *string    `gorm:"type:char(36);null;column:ModEmployee"`
}

type EmployeePhones []EmployeePhones

func (EmployeePhone) TableName() string {
	return "EmployeesPhones"
}

func (e *EmployeePhone) BeforeCreate(tx *gorm.DB) (err error) {
	e.CreationDate = time.Now()

	return
}
