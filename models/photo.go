package models

type Photo struct {
	Id       int64  `gorm:"primaryKey;not null;uniqueIndex;" json:"id"`
	Title    string `gorm:"type:varchar(300)" json:"title"`
	Caption  string `gorm:"type:varchar(300)" json:"caption"`
	PhotoUrl string `gorm:"type:varchar(300)" json:"photoUrl"`
	UserID   uint   `json:"UserID"`
}
