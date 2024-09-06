package tables

import "time"

// Table of the database
type Phone struct {
	Id               string     `gorm:"primaryKey;type:varchar(36);column:Id"`
	PhoneNumber      string     `gorm:"type:varchar(20);not null;column:PhoneNumber"`
	CreationDate     time.Time  `gorm:"type:datetime;not null;column:CreationDate"`
	CreationEmployee string     `gorm:"type:char(36);not null;column:CreationEmployee"`
	ModDate          *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee      *string    `gorm:"type:char(36);null;column:ModEmployee"`
}

func (Phone) TableName() string {
	return "Phones"
}
