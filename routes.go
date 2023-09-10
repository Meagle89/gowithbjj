package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetAllTechniques(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var techniques []Technique
	result := db.Find(&techniques)

	if result.Error != nil {
		fmt.Println("Error retrieving techniques: ", result.Error)
		http.Error(w, "Error retrieving techniques", http.StatusInternalServerError)
		return
	}

	jsonResonse, err := json.Marshal(techniques)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResonse)

}

func AddTechnique(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var newTechnique Technique
	fmt.Println("Before Decoding: ", newTechnique) // Debug
	if err := json.NewDecoder(r.Body).Decode(&newTechnique); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("After Decoding: ", newTechnique) // Debug

	if err := db.Create(&newTechnique).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.DefaultMaxHeaderBytes)
	json.NewEncoder(w).Encode(newTechnique)
}

func UpdateTechnique(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	id := mux.Vars(r)["id"]

	var updatedTechnique Technique
	if err := json.NewDecoder(r.Body).Decode(&updatedTechnique); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Model(&Technique{}).Where("id = ?", id).Updates(updatedTechnique).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTechnique)
}

func DeleteTechnique(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	id := mux.Vars(r)["id"]

	var deletedTechnique Technique
	if err := db.First(&deletedTechnique, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := db.Where("id = ?", id).Delete(&deletedTechnique).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deletedTechnique)
}
