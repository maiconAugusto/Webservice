package controller

import (
	"app/src/app/model"
	"app/src/crypto"
	"app/src/database"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Connection()
	var genericError model.ErrorMessage

	if erro != nil {
		genericError.Error = erro.Error()
		genericError.Message = "Erro"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(genericError)
		return
	}

	defer db.Close()

	response, erro := db.Query("select * from user")

	if erro != nil {
		w.WriteHeader(400)
		genericError.Error = erro.Error()
		genericError.Message = "Erro ao realizar query"
		return
	}

	defer response.Close()

	var users []model.User

	for response.Next() {
		var user model.User
		if erro := response.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created); erro != nil {
			w.WriteHeader(400)
			genericError.Error = erro.Error()
			genericError.Message = "Erro ao realizar o SCAN"
			return
		}
		users = append(users, user)
	}

	w.WriteHeader(200)

	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.WriteHeader(400)
		genericError.Error = erro.Error()
		genericError.Message = "Erro ao realizar o parser"
		return
	}

}
func Show(w http.ResponseWriter, r *http.Request) {
	var genericError model.ErrorMessage
	params := mux.Vars(r)
	ID, erro := strconv.ParseUint(params["id"], 10, 32)

	if erro != nil {
		genericError.Error = erro.Error()
		genericError.Message = "Erro ao obter ID"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(genericError)
		return
	}

	db, erro := database.Connection()

	if erro != nil {
		genericError.Error = erro.Error()
		genericError.Message = "Erro ao realizar ao conectar com o banco de dados"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(genericError)
		return
	}

	defer db.Close()

	response, erro := db.Query("select * from user where id = ?", ID)

	if erro != nil {
		genericError.Error = erro.Error()
		genericError.Message = "Erro ao realizar ao realizar prepare"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(genericError)
		return
	}

	defer response.Close()

	var user model.User

	if response.Next() {
		if erro := response.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created); erro != nil {
			genericError.Error = erro.Error()
			genericError.Message = "Erro ao realizar inferencia de dados"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(genericError)
			return
		}
	}

	if erro := json.NewEncoder(w).Encode(user); erro != nil {
		genericError.Error = erro.Error()
		genericError.Message = "Erro ao realizar o parser da query"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(genericError)
		return
	}
}
func Store(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var genericError model.ErrorMessage

	body, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		genericError.Error = erro.Error()
		genericError.Message = "Erro ao ler body"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(genericError)
		return
	}

	if erro := json.Unmarshal(body, &user); erro != nil {
		genericError.Error = erro.Error()
		genericError.Message = "Error, body empty"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(genericError)
		return
	}

	db, erro := database.Connection()

	if erro != nil {
		genericError.Error = erro.Error()
		genericError.Message = "Error"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(genericError)
		return
	}

	defer db.Close()
	db.Exec("dev")

	statement, erro := db.Prepare("insert into user (name, email, password) values (?, ?, ?)")

	if erro != nil {
		genericError.Error = erro.Error()
		genericError.Message = "Error"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(genericError)
		return
	}
	password, _ := crypto.CreateHash(user.Password)
	response, erro := statement.Exec(user.Name, user.Email, password)

	if erro != nil {
		genericError.Error = erro.Error()
		genericError.Message = "Error"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(genericError)
		return
	}

	statement.Close()

	Id, _ := response.LastInsertId()
	user.Id = uint64(Id)
	user.Password = string(password)
	json.NewEncoder(w).Encode(user)

}
func Update(w http.ResponseWriter, r *http.Request) {

}
func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var genericError model.ErrorMessage

	ID, erro := strconv.ParseUint(params["id"], 10, 32)

	if erro != nil {
		genericError.Error = erro.Error()
		genericError.Message = "Error"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(genericError)
		return
	}

	db, erro := database.Connection()

	if erro != nil {
		genericError.Error = erro.Error()
		genericError.Message = "Error"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(genericError)
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("delete from user where id = ?")

	if erro != nil {
		genericError.Error = erro.Error()
		genericError.Message = "Error"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(genericError)
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		genericError.Error = erro.Error()
		genericError.Message = "Error ao deletar usuário"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(genericError)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Deletado com sucesso")

}
