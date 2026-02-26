package model

type Job struct {
	ID     int     `gorm:"primaryKey" json:"id"`
	Title  string  `json:"title"`
	Salary float64 `json:"salary"`
}
