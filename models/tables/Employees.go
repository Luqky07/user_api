package tables

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Table of the database
type Employee struct {
	Id               string     `gorm:"primaryKey;type:varchar(36);not null;column:Id"`
	Name             string     `gorm:"type:varchar(50);not null;column:Name"`
	LastName         string     `gorm:"type:varchar(100);not null;column:LastName"`
	DocumentId       string     `gorm:"type:varchar(20);not null;column:DocumentId"`
	Email            string     `gorm:"type:varchar(50);not null;column:Email"`
	RoleId           int32      `gorm:"type:int;not null;column:RoleId"`
	CreationDate     time.Time  `gorm:"type:datetime;not null;column:CreationDate"`
	CreationEmployee string     `gorm:"type:char(36);not null;column:CreationEmployee"`
	ModDate          *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee      *string    `gorm:"type:char(36);null;column:ModEmployee"`
}

type Employees []Employee

func (Employee) TableName() string {
	return "Employees"
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New().String()
	e.CreationDate = time.Now()

	return
}
