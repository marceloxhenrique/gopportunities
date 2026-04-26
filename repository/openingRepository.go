package repository

import (
	"github.com/marceloxhenrique/gopportunities/schemas"
	"gorm.io/gorm"
)

type OpeningRepository interface {
	Create(opening *schemas.Opening) error
	GetById(id uint) (*schemas.Opening, error)
	List() ([]schemas.Opening, error)
	Update(opening *schemas.Opening) (*schemas.Opening, error)
	Delete(id uint) error
}

type GormOpeningRepository struct {
	db *gorm.DB
}

func NewGormOpenRepository(db *gorm.DB) OpeningRepository {
	return &GormOpeningRepository{
		db: db,
	}
}

func (r *GormOpeningRepository) Create(opening *schemas.Opening) error {
	return r.db.Create(opening).Error
}

func (r *GormOpeningRepository) GetById(id uint) (*schemas.Opening, error) {
	var opening schemas.Opening
	err := r.db.First(&opening, id).Error
	if err != nil {
		return nil, err
	}
	return &opening, nil
}

func (r *GormOpeningRepository) List() ([]schemas.Opening, error) {
	var openings []schemas.Opening
	err := r.db.Find(&openings).Error
	return openings, err
}

func (r *GormOpeningRepository) Update(opening *schemas.Opening) (*schemas.Opening, error) {
	err := r.db.Save(opening).Error
	if err != nil {
		return nil, err
	}
	return opening, nil
}

func (r *GormOpeningRepository) Delete(id uint) error {
	return r.db.Delete(&schemas.Opening{}, id).Error
}
