package entity

// Comment Entity
type Comment struct {
	GormModel
	UserID  uint   `gorm:"NOT NULL"`
	BlogID  uint   `gorm:"NOT NULL"`
	Content string `gorm:"NOT NULL"`
	User    User   `gorm:"constraint:OnDelete:CASCADE"`
	Blog    Blog   `gorm:"constraint:OnDelete:CASCADE"`
}
