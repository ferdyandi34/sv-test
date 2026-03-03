package models

import "time"

type Article struct {
    ID        uint      `gorm:"primaryKey"`
    Title     string    `gorm:"type:varchar(200);not null"`
    Content   string    `gorm:"type:text;not null"`
    Category    string    `gorm:"type:varchar(100);not null"`
    Created_date time.Time
    Updated_date time.Time
	Status  string  `gorm:"type:enum('Publish','Draft','Trash');not null"`
}