package model

type Task struct {
	Id         int    `gorm:"type:int;primary_key"`
	Title      string `gorm:"not null" json:"title,omitempty"`
	DueDate    string `gorm:"not null" json:"due_date,omitempty"`
	Priority   int    `gorm:"not null" json:"priority,omitempty"`
	Completed  bool   `gorm:"not null" json:"completed,omitempty"`
	Items      []Item `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items,omitempty"`
	CreatedBy  int    `gorm:"not null" json:"created_by,omitempty"`
	CreatedAt  int64  `gorm:"not null" json:"created_at,omitempty"`
	ModifiedBy int    `gorm:"not null" json:"modified_by,omitempty"`
	ModifiedAt int64  `gorm:"not null" json:"modified_at,omitempty"`
}
