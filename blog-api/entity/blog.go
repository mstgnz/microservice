package entity

// Blog Entity
type Blog struct {
	GormModel
	UserID    uint   `gorm:"NOT NULL"`
	Title     string `gorm:"NOT NULL;size:255"`
	Slug      string `gorm:"NOT NULL;size:255;index:blogs_slug,unique"`
	ShortText string `gorm:"NOT NULL;size:500"`
	LongText  string `gorm:"NOT NULL"`
	User      User   `gorm:"constraint:OnDelete:CASCADE"`
}
