package tables

import (
	"time"

	"gorm.io/gorm"
)

// Table of the database
type ClientPhone struct {
	ClientId         string     `gorm:"primaryKey;type:char(36);column:ClientId"`
	PhoneId          string     `gorm:"primaryKey;type:char(36);column:PhoneId"`
	Notes            *string    `gorm:"type:varchar(100);null;column:Notes"`
	PhoneTypeId      int32      `gorm:"type:int;not null;column:PhoneTypeId"`
	CreationDate     time.Time  `gorm:"type:datetime;not null;column:CreationDate"`
	CreationEmployee string     `gorm:"type:char(36);not null;column:CreationEmployee"`
	ModDate          *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee      *string    `gorm:"type:char(36);null;column:ModEmployee"`
}

type ClientPhones []ClientPhone

func (ClientPhone) TableName() string {
	return "ClientsPhones"
}

func (e *ClientPhone) BeforeCreate(tx *gorm.DB) (err error) {
	e.CreationDate = time.Now()

	return
}
