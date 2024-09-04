package tables

import "time"

type Employee struct {
	EmployeeID       string     `gorm:"primaryKey;type:varchar(36);not null;column:EmployeeId"`
	Name             string     `gorm:"type:varchar(50);not null;column:Name"`
	LastName         string     `gorm:"type:varchar(100);not null;column:LastName"`
	DocumentID       string     `gorm:"type:varchar(20);not null;column:DocumentId"`
	Email            string     `gorm:"type:varchar(50);not null;column:Email"`
	RoleID           string     `gorm:"type:varchar(36);not null;column:RoleId"`
	CreationEmployee string     `gorm:"type:varchar(36);not null;column:CreationEmployee"`
	ModDate          *time.Time `gorm:"type:datetime;null;column:ModDate"`
	ModEmployee      *string    `gorm:"type:char(36);null;column:ModEmployee"`

	// Define la relación con EmployeesRoles
	Role EmployeeRole `gorm:"foreignKey:RoleID;references:RoleID"`
}

func (Employee) TableName() string {
	return "Employees"
}
