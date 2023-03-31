package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/ditobayu/task-5-vix-btpns-Dito-Bayu-Adhitya/helpers"
	"github.com/ditobayu/task-5-vix-btpns-Dito-Bayu-Adhitya/models"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

var JWT_KEY = []byte("ashdjqy9283409bsdklkg8hda01")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}

func Login(w http.ResponseWriter, r *http.Request) {

	// mengambil inputan json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// ambil data user berdasarkan username
	var user models.User
	if err := models.DB.Where("email = ?", userInput.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Email atau password salah"}
			helpers.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helpers.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// cek apakah password valid
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Username atau password salah"}
		helpers.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// proses pembuatan token jwt
	expTime := time.Now().Add(time.Minute * 60)
	claims := &JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// medeklarasikan algoritma yang akan digunakan untuk signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signed token
	token, err := tokenAlgo.SignedString(JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// set token yang ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "login berhasil"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {

	// mengambil inputan json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// hash pass menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert ke database
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// var user models.User
	params := mux.Vars(r)
	id:= params["userId"]

	
	// mengambil inputan json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	
	// hash pass menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	if models.DB.Model(&userInput).Where("id = ?", id).Updates(&userInput).RowsAffected == 0 {
		helpers.ResponseJSON(w, http.StatusBadRequest, "Tidak dapat menghapus product")
		return
	}
	
	response := map[string]string{"message": "change user data success"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id:= params["userId"]
	
	// mengambil inputan json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	
	defer r.Body.Close()
	
	if models.DB.Model(&userInput).Where("id = ?", id).Delete(&userInput).RowsAffected == 0 {
		helpers.ResponseJSON(w, http.StatusBadRequest, "Tidak dapat menghapus product")
		return
	}


	response := map[string]string{"message": "delete user success"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}
