package registrations

import (
	"gorm.io/gorm"
	"time"
  "github.com/TCSOfficial/RegistrationsBackend/database"
)

var db *gorm.DB
func Init() {
	db = database.DB
	db.AutoMigrate(&Registration{})
}

type Registration struct {
  ID          uint           `gorm:"primaryKey"`
  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `json:"name"`
	CollegeID   string         `json:"college_id"`
	Email       string         `json:"e_mail"`
	Phone       string         `json:"phone_number"`
}
