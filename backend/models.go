package main

type Technique struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Type     string
	Position string
}
