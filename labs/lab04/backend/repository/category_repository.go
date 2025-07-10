package repository

import (
	"lab04-backend/models"

	"gorm.io/gorm"
	// removed extra close paren
)

// CategoryRepository handles database operations for categories using GORM
// This repository demonstrates GORM ORM approach for database operations
type CategoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates a new CategoryRepository with GORM
func NewCategoryRepository(gormDB *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: gormDB}
}

// TODO: Implement Create method using GORM
func (r *CategoryRepository) Create(category *models.Category) error {
	result := r.db.Create(category)
	return result.Error
}

// TODO: Implement GetByID method using GORM
func (r *CategoryRepository) GetByID(id uint) (*models.Category, error) {
	var category models.Category
	result := r.db.First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

// TODO: Implement GetAll method using GORM
func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	result := r.db.Order("name").Find(&categories)
	return categories, result.Error
}

// TODO: Implement Update method using GORM
func (r *CategoryRepository) Update(category *models.Category) error {
	result := r.db.Save(category)
	return result.Error
}

// TODO: Implement Delete method using GORM
func (r *CategoryRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Category{}, id)
	return result.Error
}

// TODO: Implement FindByName method using GORM
func (r *CategoryRepository) FindByName(name string) (*models.Category, error) {
	var category models.Category
	result := r.db.Where("name = ?", name).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

// TODO: Implement SearchCategories method using GORM
func (r *CategoryRepository) SearchCategories(query string, limit int) ([]models.Category, error) {
	var categories []models.Category
	result := r.db.Where("name LIKE ?", "%"+query+"%").Order("name").Limit(limit).Find(&categories)
	return categories, result.Error
}

// TODO: Implement GetCategoriesWithPosts method using GORM associations
func (r *CategoryRepository) GetCategoriesWithPosts() ([]models.Category, error) {
	var categories []models.Category
	result := r.db.Preload("Posts").Order("name").Find(&categories)
	return categories, result.Error
}

// TODO: Implement Count method using GORM
func (r *CategoryRepository) Count() (int64, error) {
	var count int64
	result := r.db.Model(&models.Category{}).Count(&count)
	return count, result.Error
}

// TODO: Implement Transaction example using GORM
func (r *CategoryRepository) CreateWithTransaction(categories []models.Category) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, category := range categories {
			if err := tx.Create(&category).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
