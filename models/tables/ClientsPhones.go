package tables

import "time"

type ClientPhone struct {
	ClientId         string     `gorm:"primaryKey;type:varchar(36);column:ClientId"`
	PhoneId          string     `gorm:"primaryKey;type:varchar(36);column:PhoneId"`
	Note             *string    `gorm:"type:varchar(100);null;column:Note"`
	PhoneTypeId      int64      `gorm:"type:int;not null;column:PhoneTypeId"`
	CreationDate     time.Time  `gorm:"type:datetime;not null;column:CreationDate"`
	CreationEmployee string     `gorm:"type:varchar(36);not null;column:CreationEmployee"`
	ModDate          *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee      *string    `gorm:"type:varchar(36);null;column:ModEmployee"`
}

func (ClientPhone) TableName() string {
	return "ClientsPhones"
}
