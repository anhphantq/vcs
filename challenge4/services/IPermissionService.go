package services

import (
	"challenge4/models"
	"challenge4/repositories"
	"errors"
)

type IPermissionService struct {
	repositories.PermissionRepository
}

func (srv *IPermissionService) GetAllPermission() ([]models.Permission, error) {
	return srv.FindPermission(models.Permission{})
}

func (srv *IPermissionService) GetPermissionByID(id uint) (models.Permission, error) {
	Permissions, err := srv.FindPermission(models.Permission{Permission_id: id})

	if err != nil {
		return models.Permission{}, err
	}

	if len(Permissions) < 1 {
		return models.Permission{}, errors.New("no Permissions founded")
	}

	return Permissions[0], nil
}

func (srv *IPermissionService) InsertPermission(Permission models.Permission) (models.Permission, error) {
	return srv.SavePermission(Permission)
}

func (srv *IPermissionService) DeletePermissionByID(id uint) error {
	return srv.DeletePermission(models.Permission{Permission_id: id})
}
