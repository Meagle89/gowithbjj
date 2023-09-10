package main

import (
	"fmt"

	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) {
	var count int64
	db.Model(&Technique{}).Count(&count)

	if count > 0 {
		fmt.Println("Database already seeded")
		return
	}

	techniques := []Technique{
		{Name: "Armbar", Type: "Submission", Position: "Guard"},
		{Name: "Triangle Choke", Type: "Submission", Position: "Guard"},
		{Name: "Kimura", Type: "Submission", Position: "Side Control"},
		{Name: "Rear Naked Choke", Type: "Submission", Position: "Back Mount"},
		{Name: "Guillotine", Type: "Submission", Position: "Guard"},
		{Name: "Sweep", Type: "Transition", Position: "Guard"},
		{Name: "Pass", Type: "Transition", Position: "Open Guard"},
	}

	for _, technique := range techniques {
		db.Create(&technique)
	}
}
