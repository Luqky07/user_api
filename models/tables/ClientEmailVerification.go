package tables

import (
	"time"

	"gorm.io/gorm"
)

// Table of the database
type ClientEmailVerification struct {
	ClientId         string     `gorm:"primaryKey;type:char(36);column:ClientId"`
	IsVerified       bool       `gorm:"type:bit;not null;column:IsVerified"`
	Token            string     `gorm:"type:char(36);not null;column:Token"`
	Expiration       time.Time  `gorm:"type:datetime;not null:column:Expiration"`
	CreationDate     time.Time  `gorm:"type:datetime;not null;column:CreationDate"`
	CreationEmployee string     `gorm:"type:char(36);not null;column:CreationEmployee"`
	ModDate          *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee      *string    `gorm:"type:char(36);null;column:ModEmployee"`
}

type ClientEmailVerifications []ClientEmailVerification

func (ClientEmailVerification) TableName() string {
	return "ClientsEmailVerification"
}

func (e *ClientEmailVerification) BeforeCreate(tx *gorm.DB) (err error) {
	e.CreationDate = time.Now()

	return
}
