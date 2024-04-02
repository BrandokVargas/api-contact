package user

import (
	"time"

	"github.com/BrandokVargas/api-contact/model"
)

//Implemetación codigo sql
//Consultas etc

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate user db
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

func (s *Service) Create(u *model.User) error {
	u.CreatedAt = time.Now()
	return s.storage.Create(u)
}
