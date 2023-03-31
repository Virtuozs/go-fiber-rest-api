package model

type Item struct {
	Id        int    `gorm:"type:int;primary_key"`
	TaskId    int    `gorm:"not null" json:"-"`
	Desc      string `gorm:"not null" json:"description,omitempty"`
	Completed bool   `gorm:"not null" json:"completed,omitempty"`
}
