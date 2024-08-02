package models

type Type struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `gorm:"type:varchar(50);not null" json:"name"`
	Code      string `gorm:"type:varchar(50);not null" json:"code"`
	Category  string `gorm:"index:idx_category;type:varchar(50);not null" json:"category"`
	CreatedAt string `gorm:"autoCreateTime" json:"createdAt"`
	UpdateAt  string `gorm:"autoUpdateTime" json:"updatedAt"`
}
