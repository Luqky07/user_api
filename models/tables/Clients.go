package tables

import "time"

type Client struct {
	Id               string     `gorm:"primaryKey;type:varchar(36);column:Id"`
	Name             string     `gorm:"type:varchar(50);not null;column:Name"`
	LastName         string     `gorm:"type:varchar(100);not null;column:LastName"`
	Birthdate        time.Time  `gorm:"type:datetime;not null;column:Birthdate"`
	Email            string     `gorm:"type:varchar(50);not null;column:Email"`
	DocumentId       string     `gorm:"type:varchar(50);not null;column:DocumentId"`
	CreationDate     time.Time  `gorm:"type:datetime;not null;column:CreationDate"`
	CreationEmployee string     `gorm:"type:varchar(36);not null;column:CreationEmployee"`
	ModDate          *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee      *string    `gorm:"type:varchar(36);null;column:ModEmployee"`
}

func (Client) TableName() string {
	return "Clients"
}
