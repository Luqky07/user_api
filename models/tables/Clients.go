package tables

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Table of the database
type Client struct {
	Id               string     `gorm:"primaryKey;type:char(36);column:Id"`
	Name             string     `gorm:"type:varchar(50);not null;column:Name"`
	LastName         string     `gorm:"type:varchar(100);not null;column:LastName"`
	Birthdate        time.Time  `gorm:"type:datetime;not null;column:Birthdate"`
	Email            string     `gorm:"type:varchar(50);not null;column:Email"`
	DocumentId       string     `gorm:"type:varchar(50);not null;column:DocumentId"`
	CreationDate     time.Time  `gorm:"type:datetime;not null;column:CreationDate"`
	CreationEmployee string     `gorm:"type:char(36);not null;column:CreationEmployee"`
	ModDate          *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee      *string    `gorm:"type:char(36);null;column:ModEmployee"`
}

type Clients []Client

func (Client) TableName() string {
	return "Clients"
}

func (e *Client) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New().String()
	e.CreationDate = time.Now()

	return
}
