package model

type Job struct {
	ID     int     `gorm:"primaryKey" json:"id" binding:"required"`
	Title  string  `json:"title"`
	Salary float64 `json:"salary"`
}
