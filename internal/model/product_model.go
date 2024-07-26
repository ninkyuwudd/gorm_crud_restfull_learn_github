package model

import "time"

type Product struct {
	ID         int       `gorm:"primary_key;column:id"`
	Name       string    `gorm:"column:name"`
	CategoryID int       `gorm:"primary_key;column:category_id"`
	Category   Category  `gorm:"foreignKey:CategoryID"`
	Quantity   int       `gorm:"column:quantity"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

type ProductUpdate struct {
	Name      string    `gorm:"column:name"`
	Category  string    `gorm:"column:category"`
	Quantity  int       `gorm:"column:quantity"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
