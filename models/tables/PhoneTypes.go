package tables

import "time"

// Table of the database
type PhoneType struct {
	Id               int64      `gorm:"primaryKey;autoincrement;type:int;column:Id"`
	Description      string     `gorm:"type:varchar(50);not null;columns:Description"`
	CreationDate     time.Time  `gorm:"type:datetime;not null;column:CreationDate"`
	CreationEmployee string     `gorm:"type:char(36);not null;column:CreationEmployee"`
	ModDate          *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee      *string    `gorm:"type:char(36);null;column:ModEmployee"`
}

func (PhoneType) TableName() string {
	return "PhonesTypes"
}
