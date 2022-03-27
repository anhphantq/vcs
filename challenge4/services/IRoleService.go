package services

import (
	"challenge4/models"
	"challenge4/repositories"
	"errors"
)

type IRoleService struct {
	repositories.RoleRepository
}

func (srv *IRoleService) GetAllRole() ([]models.Role, error) {
	roles, err := srv.FindRole(models.Role{})

	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (srv *IRoleService) GetRoleByID(id uint) (models.Role, error) {
	roles, err := srv.FindRole(models.Role{Role_id: id})

	if err != nil {
		return models.Role{}, err
	}

	if len(roles) < 1 {
		return models.Role{}, errors.New("no roles founded")
	}

	return roles[0], nil
}

func (srv *IRoleService) InsertRole(role models.Role) (models.Role, error) {
	return srv.SaveRole(role)
}

func (srv *IRoleService) DeleteRoleByID(id uint) error {
	return srv.DeleteRole(models.Role{Role_id: id})
}
