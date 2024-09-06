package tables

import "time"

// Table of the database
type ClientCredential struct {
	ClientId         string     `gorm:"primaryKey;type:varchar(36);column:ClientId"`
	IsBlock          bool       `gorm:"type:bit;not null;column:IsBlock"`
	FailAttempts     int16      `gorm:"type:smallint;not null;column:FailAttempts"`
	LastLogin        time.Time  `gorm:"type:datetime;not null;column:LastLogin"`
	Password         string     `gorm:"type:varchar(64);not null;column:Password"`
	Token            string     `gorm:"type:varchar(36);not null;column:Token"`
	Expiration       time.Time  `gorm:"type:datetime;not null:column:Expiration"`
	CreationDate     time.Time  `gorm:"type:datetime;not null;column:CreationDate"`
	CreationEmployee string     `gorm:"type:varchar(36);not null;column:CreationEmployee"`
	ModDate          *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee      *string    `gorm:"type:varchar(36);null;column:ModEmployee"`
}

func (ClientCredential) TableName() string {
	return "ClientsCredentials"
}
