package handler

import (
	"encoding/json"
	"net/http"

	"github.com/BrandokVargas/api-contact/model"
	"github.com/BrandokVargas/api-contact/services/user"
)

type User struct {
	user user.Storage
}

func NewUser(u user.Storage) *User {
	return &User{u}
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`
		{
			"message_type": "error",
			"message": "Metodo no permitido"
		}`))
		return
	}

	data := &model.User{}

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`
		{
			"message_type": "error",
			"message": "La estructura del request está mal escrita"
		}`))
		return
	}

	err = u.user.Create(data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`
		{
			"message_type": "error",
			"message": "Hubo un problema con el servidor al crear el usuario"
		}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`
	{
		"message_type": "ok",
		"message": "User creado correctamente"
	}`))

}
