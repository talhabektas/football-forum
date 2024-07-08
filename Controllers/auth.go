package controllers

import (
	"encoding/json"
	"football-forum/database"
	"football-forum/middleware"
	"football-forum/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// Register registers a new user
func Register(w http.ResponseWriter, r *http.Request) {
	log.Println("Register endpoint hit")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Şifre hashlenirken hata oluştu: ", err)
		http.Error(w, "Şifre hashlenirken hata oluştu", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	log.Println("Veritabanında kullanıcı oluşturuluyor")
	result := database.DB.Create(&user)
	if result.Error != nil {
		log.Println("Kullanıcı oluşturulurken hata oluştu: ", result.Error)
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Kullanıcı başarıyla oluşturuldu")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Login logs in a user
func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Login endpoint hit")
	var input struct {
		username string
		Email    string
		Password string
	}
	json.NewDecoder(r.Body).Decode(&input)

	var user models.User
	result := database.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		log.Println("Geçersiz e-posta veya şifre")
		http.Error(w, "Geçersiz e-posta veya şifre", http.StatusUnauthorized)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		log.Println("Geçersiz e-posta veya şifre")
		http.Error(w, "Geçersiz e-posta veya şifre", http.StatusUnauthorized)
		return
	}

	token, err := middleware.GenerateJWT(user.Email)
	if err != nil {
		log.Println("Token oluşturulurken hata: ", err)
		http.Error(w, "Token oluşturulurken hata", http.StatusInternalServerError)
		return
	}

	log.Println("Kullanıcı başarıyla giriş yaptı")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
