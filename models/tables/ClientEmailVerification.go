package tables

import "time"

// Table of the database
type ClientEmailVerification struct {
	ClientId         string     `gorm:"primaryKey;type:varchar(36);column:ClientId"`
	IsVerified       bool       `gorm:"type:bit;not null;column:IsVerified"`
	Token            string     `gorm:"type:varchar(36);not null;column:Token"`
	Expiration       time.Time  `gorm:"type:datetime;not null:column:Expiration"`
	CreationDate     time.Time  `gorm:"type:datetime;not null;column:CreationDate"`
	CreationEmployee string     `gorm:"type:varchar(36);not null;column:CreationEmployee"`
	ModDate          *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee      *string    `gorm:"type:varchar(36);null;column:ModEmployee"`
}

func (ClientEmailVerification) TableName() string {
	return "ClientsEmailVerification"
}
