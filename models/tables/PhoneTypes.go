package tables

import (
	"time"

	"gorm.io/gorm"
)

// Table of the database
type PhoneType struct {
	Id               int64      `gorm:"primaryKey;autoincrement;type:int;column:Id"`
	Description      string     `gorm:"type:varchar(50);not null;column:Description"`
	CreationDate     time.Time  `gorm:"type:datetime;not null;column:CreationDate"`
	CreationEmployee string     `gorm:"type:char(36);not null;column:CreationEmployee"`
	ModDate          *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee      *string    `gorm:"type:char(36);null;column:ModEmployee"`
}

type PhoneTypes []PhoneType

func (PhoneType) TableName() string {
	return "PhoneTypes"
}

func (e *PhoneType) BeforeCreate(tx *gorm.DB) (err error) {
	e.CreationDate = time.Now()

	return
}
