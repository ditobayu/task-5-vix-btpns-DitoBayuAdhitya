package models

// "time"

type User struct {
	Id        int64  `gorm:"primaryKey;not null;uniqueIndex;" json:"id"`
	Username  string `gorm:"type:varchar(300)" json:"username"`
	Email     string `gorm:"type:varchar(300)" json:"email"`
	Password  string `gorm:"type:varchar(300)" json:"password"`
	Photo []Photo `gorm:"foreignKey:id" json:"photo"`
	CreatedAt []uint8
	UpdatedAt []uint8
}