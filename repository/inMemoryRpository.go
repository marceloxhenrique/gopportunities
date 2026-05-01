package repository

import (
	"errors"

	"github.com/marceloxhenrique/gopportunities/schemas"
)

type InMemoryOpeningRepository struct {
	Data   map[uint]*schemas.Opening
	NextID uint
}

func NewInMemoryRepository() *InMemoryOpeningRepository {
	return &InMemoryOpeningRepository{
		Data: make(map[uint]*schemas.Opening),
	}
}

func (r *InMemoryOpeningRepository) Create(opening *schemas.Opening) error {
	r.NextID++
	opening.ID = r.NextID
	r.Data[opening.ID] = opening
	return nil
}

func (r *InMemoryOpeningRepository) GetById(id uint) (*schemas.Opening, error) {
	if o, ok := r.Data[id]; ok {
		return o, nil
	}
	return nil, errors.New("not found")
}

func (r *InMemoryOpeningRepository) List() ([]schemas.Opening, error) {
	return nil, nil
}
func (r *InMemoryOpeningRepository) Update(opening *schemas.Opening) (*schemas.Opening, error) {
	return nil, nil
}
func (r *InMemoryOpeningRepository) Delete(id uint) error {
	if _, ok := r.Data[id]; !ok {
		return errors.New("not found")
	}
	delete(r.Data, id)
	return nil
}
