package tables

import "time"

type Client struct {
	ClientID            string     `gorm:"primaryKey;type:varchar(36);column:ClientId"`
	Name                string     `gorm:"type:varchar(50);not null;column:Name"`
	LastName            string     `gorm:"type:varchar(100);not null;column:LastName"`
	Birthdate           time.Time  `gorm:"type:datetime;not null;column:Birthdate"`
	Email               string     `gorm:"type:varchar(50);not null;column:Email"`
	DocumentID          string     `gorm:"type:varchar(50);not null;column:DocumentId"`
	CreationDate        time.Time  `gorm:"type:datetime;not null;column:CreationDate"`
	CreationEmployee    string     `gorm:"type:varchar(36);not null;column:CreationEmployee"`
	CreationEmployeeRef Employee   `gorm:"foreignKey:CreationEmployee;references:EmployeeID"`
	ModDate             *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee         *string    `gorm:"type:varchar(36);null;column:ModEmployee"`
	ModEmployeeRef      Employee   `gorm:"foreignKey:ModEmployee;references:EmployeeID"`
}

func (Client) TableName() string {
	return "Clients"
}
