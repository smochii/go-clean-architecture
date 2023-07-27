package database

import (
	"github.com/samber/lo"
	"github.com/smochii/go-clean-architecture/domain"
	"github.com/smochii/go-clean-architecture/domain/entity"
	"github.com/smochii/go-clean-architecture/domain/value"
	"github.com/smochii/go-clean-architecture/infrastructure/database/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user entity.User) (*entity.User, error) {
	model := r.entityToModel(user)
	err := r.db.Create(&model).Error
	if err != nil {
		return nil, err
	}
	entity := r.modelToEntity(model)
	return &entity, err
}

func (r *UserRepository) Update(user entity.User) (*entity.User, error) {
	model := r.entityToModel(user)
	err := r.db.Updates(&model).Error
	if err != nil {
		return nil, err
	}
	entity := r.modelToEntity(model)
	return &entity, nil
}

func (r *UserRepository) FindById(id value.UserId) (*entity.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	entity := r.modelToEntity(user)
	return &entity, nil
}

func (r *UserRepository) FindByEmail(email value.Email) (*entity.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	entity := r.modelToEntity(user)
	return &entity, nil
}

func (r *UserRepository) FindAll() ([]entity.User, error) {
	var models []model.User
	err := r.db.Find(&models).Error
	if err != nil {
		return nil, err
	}
	entities := lo.Map(models, func(m model.User, i int) entity.User {
		return r.modelToEntity(m)
	})
	return entities, nil
}

func (r *UserRepository) modelToEntity(model model.User) entity.User {
	return entity.NewUser(
		model.Id,
		model.Email,
		model.Password,
		model.CreatedAt,
		model.UpdatedAt,
	)
}

func (r *UserRepository) entityToModel(entity entity.User) model.User {
	return model.User{
		Id:        entity.Id(),
		Email:     entity.Email(),
		Password:  entity.Password(),
		CreatedAt: entity.CreatedAt(),
		UpdatedAt: entity.UpdatedAt(),
	}
}
