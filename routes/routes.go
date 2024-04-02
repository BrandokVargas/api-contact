package routes

import (
	"net/http"

	"github.com/BrandokVargas/api-contact/handler"
	"github.com/BrandokVargas/api-contact/services/user"
)

// Route User
func RouteUser(mux *http.ServeMux, user user.Storage) {
	route := handler.NewUser(user)
	mux.HandleFunc("/api/v1/user", route.Create)

}
