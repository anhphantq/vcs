package services

import (
	"challenge4/models"
	"challenge4/repositories"
	"errors"
)

type IPostService struct {
	repositories.PostRepository
}

func (srv *IPostService) GetPostByUserID(user_id uint, page int) ([]models.Post, error) {
	posts, err := srv.FindPost(models.Post{User_id: user_id}, page*10, 10)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (srv *IPostService) GetPostByID(post_id uint) (models.Post, error) {
	posts, err := srv.FindPost(models.Post{Post_id: post_id}, 0, 1)

	if err != nil {
		return models.Post{}, err
	}

	if len(posts) < 1 {
		return models.Post{}, errors.New("no posts founded")
	}

	return posts[0], err
}

func (srv *IPostService) InsertPost(post models.Post) (models.Post, error) {
	post, err := srv.SavePost(post)

	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (srv *IPostService) UpdatePostSrv(post models.Post) (models.Post, error) {
	post, err := srv.UpdatePost(post)

	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (srv *IPostService) DeletePostSrv(post_id uint) error {
	err := srv.DeletePost(models.Post{Post_id: post_id})

	if err != nil {
		return err
	}

	return nil
}

func (srv *IPostService) GetAllPosts(page int) ([]models.Post, error) {
	posts, err := srv.FindPost(models.Post{}, page*10, 10)

	if err != nil {
		return nil, err
	}

	return posts, nil
}
