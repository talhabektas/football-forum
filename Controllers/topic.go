package controllers

import (
	"encoding/json"
	"football-forum/models"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateTopic(w http.ResponseWriter, r *http.Request) {
	var topic models.Topic
	json.NewDecoder(r.Body).Decode(&topic)

	result := db.Create(&topic)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(topic)
}

func GetTopics(w http.ResponseWriter, r *http.Request) {
	var topics []models.Topic
	db.Preload("User").Find(&topics)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(topics)
}

func DeleteTopic(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result := db.Delete(&models.Topic{}, params["id"])
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
