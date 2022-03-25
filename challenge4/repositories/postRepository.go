package repositories

import "challenge3/models"

type PostRepository interface {
	SavePost(post models.Post) (models.Post, error)
	FindPost(post models.Post) ([]models.Post, error)
	DeletePostByID(id uint) error
	UpdatePostByID(id uint, post models.Post) (models.Post, error)
}

type IPostRepository struct {
}
