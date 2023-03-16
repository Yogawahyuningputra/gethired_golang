package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type TodosRepository interface {
	FindTodos() ([]models.Todos, error)
	GetTodos(ID int) (models.Todos, error)
	CreateTodos(todos models.Todos) (models.Todos, error)
	UpdateTodos(todos models.Todos) (models.Todos, error)
	DeleteTodos(todos models.Todos) (models.Todos, error)
}

func RepositoryTodos(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTodos() ([]models.Todos, error) {
	var todos []models.Todos
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *repository) GetTodos(ID int) (models.Todos, error) {
	var todo models.Todos
	err := r.db.First(&todo, ID).Error
	return todo, err
}
func (r *repository) CreateTodos(todo models.Todos) (models.Todos, error) {
	err := r.db.Create(&todo).Error
	return todo, err
}

func (r *repository) UpdateTodos(todo models.Todos) (models.Todos, error) {
	err := r.db.Save(&todo).Error
	return todo, err
}

func (r *repository) DeleteTodos(todo models.Todos) (models.Todos, error) {
	err := r.db.Delete(&todo).Error
	return todo, err
}
