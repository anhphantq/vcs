package repositories

import (
	"challenge4/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	SavePost(post models.Post) (models.Post, error)
	FindPost(post models.Post, offset int, limit int) ([]models.Post, error)
	DeletePost(post models.Post) error
	UpdatePost(post models.Post) (models.Post, error)
}

type IPostRepository struct {
	DB *gorm.DB
}

func (repo *IPostRepository) SavePost(post models.Post) (models.Post, error) {
	result := repo.DB.Model(&models.Post{}).Create(&post)

	if result.Error != nil {
		return models.Post{}, result.Error
	}

	return post, nil
}
func (repo *IPostRepository) FindPost(post models.Post, offset int, limit int) ([]models.Post, error) {
	var posts []models.Post

	result := repo.DB.Model(&models.Post{}).Offset(offset).Limit(limit).Where(&post).Scan(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}
func (repo *IPostRepository) DeletePost(post models.Post) error {
	result := repo.DB.Model(&models.Post{}).Delete(&post)

	return result.Error
}
func (repo *IPostRepository) UpdatePost(post models.Post) (models.Post, error) {
	result := repo.DB.Save(&post)

	if result.Error != nil {
		return models.Post{}, result.Error
	}

	return post, nil
}
