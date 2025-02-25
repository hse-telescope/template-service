package facade

import (
	"context"

	"github.com/hse-telesope/template-service/internal/repository/models"
)

type Storage interface {
	GetPeople(ctx context.Context) ([]models.Person, error)
}

type Facade struct {
	storage Storage
}

func New(storage Storage) Facade {
	return Facade{
		storage: storage,
	}
}

func (f Facade) GetPeople(ctx context.Context) ([]models.Person, error) {
	return f.storage.GetPeople(ctx)
}
