package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type ActivityRepository interface {
	FindActivity() ([]models.Activity, error)
	GetActivity(ID int) (models.Activity, error)
	CreateActivity(activity models.Activity) (models.Activity, error)
	UpdateActivity(activity models.Activity) (models.Activity, error)
	DeleteActivity(activity models.Activity) (models.Activity, error)
}

func RepositoryActivity(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindActivity() ([]models.Activity, error) {
	var activities []models.Activity
	err := r.db.Find(&activities).Error
	return activities, err
}

func (r *repository) GetActivity(ID int) (models.Activity, error) {
	var activity models.Activity
	err := r.db.First(&activity, ID).Error
	return activity, err
}
func (r *repository) CreateActivity(activity models.Activity) (models.Activity, error) {
	err := r.db.Create(&activity).Error
	return activity, err
}

func (r *repository) UpdateActivity(activity models.Activity) (models.Activity, error) {
	err := r.db.Save(&activity).Error
	return activity, err
}

func (r *repository) DeleteActivity(activity models.Activity) (models.Activity, error) {
	err := r.db.Delete(&activity).Error
	return activity, err
}
