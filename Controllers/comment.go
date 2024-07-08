package controllers

import (
	"encoding/json"
	"football-forum/models"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	json.NewDecoder(r.Body).Decode(&comment)

	result := db.Create(&comment)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var comments []models.Comment
	db.Where("topic_id = ?", params["id"]).Preload("User").Find(&comments)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result := db.Delete(&models.Comment{}, params["id"])
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
