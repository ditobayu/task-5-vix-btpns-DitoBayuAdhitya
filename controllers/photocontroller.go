package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ditobayu/task-5-vix-btpns-Dito-Bayu-Adhitya/helpers"
	"github.com/ditobayu/task-5-vix-btpns-Dito-Bayu-Adhitya/models"

	// "gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

func GetPhoto(w http.ResponseWriter, r *http.Request) {

	// mengambil inputan json
	var userInput models.User
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	
	if err := models.DB.Where("email = ?", userInput.Email).First(&user).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// hash pass menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	if models.DB.Model(&userInput).Where("id = ?", user.Id).Updates(&userInput).RowsAffected == 0 {
		helpers.ResponseJSON(w, http.StatusBadRequest, "Tidak dapat menghapus product")
		return
	}

	// response := map[string]string{"message": "login berhasil"}
	helpers.ResponseJSON(w, http.StatusOK, user)
}
func AddPhoto(w http.ResponseWriter, r *http.Request) {

	// mengambil inputan json
	var userInput models.Photo
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	
	if err := models.DB.Model(&user).Where("id = ?", userInput.UserID).First(&user).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	newPhoto := models.Photo{
		Id: userInput.Id,
		Title    : userInput.Title,
		Caption  : userInput.Caption,
		PhotoUrl : userInput.PhotoUrl,
		UserID   : userInput.UserID,
	}
	asd := []models.Photo{
		
	}

		asd = append(asd, newPhoto)
		// asd = append(asd, newPhoto)
	user.Photo = asd
	user.Username = "popo"

	if models.DB.Model(&user).Where("id = ?", userInput.UserID).Updates(&user).RowsAffected == 0 {
		helpers.ResponseJSON(w, http.StatusBadRequest, "Tidak dapat menambah foto")
		return
	}	

	// response := map[string]string{"message": "success"}
	helpers.ResponseJSON(w, http.StatusOK, user)
}
func UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	// var user models.Photo
	params := mux.Vars(r)
	id:= params["userId"]

	
	// mengambil inputan json
	var userInput models.Photo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if models.DB.Model(&userInput).Where("id = ?", id).Updates(&userInput).RowsAffected == 0 {
		helpers.ResponseJSON(w, http.StatusBadRequest, "Tidak dapat menghapus product")
		return
	}
	
	response := map[string]string{"message": "change user data success"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}
func DeletePhoto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id:= params["userId"]
	
	// mengambil inputan json
	var userInput models.Photo
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
