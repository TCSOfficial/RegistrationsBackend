package registrations

import (
	"gorm.io/gorm"
	"time"
  "github.com/TCSOfficial/RegistrationsBackend/database"
)

var db *gorm.DB
func Init() {
	db = database.DB
	db.AutoMigrate(&Registration{}, &RecruitmentRegistration{})
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

type RecruitmentRegistration struct {
  ID          uint           `gorm:"primaryKey"`
  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `gorm:"not null" json:"name"`
	CollegeID   string         `gorm:"not null" json:"college_id"`
	Email       string         `gorm:"unique;not null" json:"e_mail"`
	Phone       string         `gorm:"unique;not null" json:"phone_number"`
}
