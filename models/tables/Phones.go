package tables

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Table of the database
type Phone struct {
	Id               string     `gorm:"primaryKey;type:char(36);column:Id"`
	PhoneNumber      string     `gorm:"type:varchar(20);not null;column:PhoneNumber"`
	CreationDate     time.Time  `gorm:"type:datetime;not null;column:CreationDate"`
	CreationEmployee string     `gorm:"type:char(36);not null;column:CreationEmployee"`
	ModDate          *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee      *string    `gorm:"type:char(36);null;column:ModEmployee"`
}

type Phones []Phone

func (Phone) TableName() string {
	return "Phones"
}

func (e *Phone) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New().String()
	e.CreationDate = time.Now()

	return
}
