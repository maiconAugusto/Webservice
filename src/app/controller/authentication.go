package controller

import (
	"app/src/app/model"
	"app/src/authentication"
	"app/src/crypto"
	"app/src/database"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {

	var errorGeneric model.Error
	var user model.User

	body, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		errorGeneric.Error = erro.Error()
		errorGeneric.Message = "Error"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errorGeneric)
		return
	}

	if erro := json.Unmarshal(body, &user); erro != nil {
		errorGeneric.Error = erro.Error()
		errorGeneric.Message = "Error"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errorGeneric)
		return
	}

	db, erro := database.Connection()

	if erro := json.Unmarshal(body, &user); erro != nil {
		errorGeneric.Error = erro.Error()
		errorGeneric.Message = "Error ao se conectar com o banco de dados"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errorGeneric)
		return
	}

	defer db.Close()

	response, erro := db.Query("select id, email, password from user where email = ?", user.Email)

	if erro := json.Unmarshal(body, &user); erro != nil {
		errorGeneric.Error = erro.Error()
		errorGeneric.Message = "Error ao se conectar com o banco de dados"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errorGeneric)
		return
	}

	defer response.Close()

	var userReponse model.User

	if response.Next() {
		if erro := response.Scan(&userReponse.Id, &userReponse.Email, &userReponse.Password); erro != nil {
			errorGeneric.Error = erro.Error()
			errorGeneric.Message = "Error ao realizar o SCAN"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errorGeneric)
			return
		}
	}

	if userReponse.Email == "" {
		errorGeneric.Error = "not found"
		errorGeneric.Message = "E-mail não encontrado"
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(errorGeneric)
		return
	}

	validate := crypto.VerifyHash(userReponse.Password, user.Password)

	token, erro := authentication.GenerateToken(user.Id)

	if erro != nil {
		errorGeneric.Error = erro.Error()
		errorGeneric.Message = "Error ao gerar token"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errorGeneric)
		return
	}

	userReponse.Token = token

	if validate != nil {
		errorGeneric.Error = validate.Error()
		errorGeneric.Message = "Senha incorreta"
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(errorGeneric)
		return
	}

	w.WriteHeader(200)
	if erro := json.NewEncoder(w).Encode(userReponse); erro != nil {
		errorGeneric.Error = erro.Error()
		errorGeneric.Message = "Error ao realizar o parser"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errorGeneric)
		return
	}
}
