package user

import "github.com/BrandokVargas/api-contact/model"

//Metodos que va a tener tu implementación
//de user
type Storage interface {
	Migrate() error
	Create(u *model.User) error
}
