package tables

import (
	"time"

	"gorm.io/gorm"
)

// Table of the database
type EmployeeCredential struct {
	EmployeeId       string     `gorm:"primaryKey;type:char(36);column:EmployeeId"`
	IsBlock          bool       `gorm:"type:bit;not null;column:IsBlock"`
	FailAttempts     int16      `gorm:"type:smallint;not null;column:FailAttempts"`
	LastLogin        time.Time  `gorm:"type:datetime;not null;column:LastLogin"`
	Password         string     `gorm:"type:varchar(64);not null;column:Password"`
	CreationDate     time.Time  `gorm:"type:datetime;not null;column:CreationDate"`
	CreationEmployee string     `gorm:"type:char(36);not null;column:CreationEmployee"`
	ModDate          *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee      *string    `gorm:"type:char(36);null;column:ModEmployee"`
}

type EmployeeCredentials []EmployeeCredential

func (EmployeeCredential) TableName() string {
	return "EmployeesCredentials"
}

func (e *EmployeeCredential) BeforeCreate(tx *gorm.DB) (err error) {
	e.CreationDate = time.Now()

	return
}
