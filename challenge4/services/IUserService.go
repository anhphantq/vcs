package services

import (
	"challenge4/models"
	"challenge4/repositories"
	"errors"
)

type IUserSerivce struct {
	repositories.UserRepository
	repositories.RoleRepository
	repositories.PermissionRepository
	repositories.GrantRepository
}

func (srv *IUserSerivce) CheckEmailUsed(email string) (bool, error) {
	users, err := srv.FindUser(models.Account{Email: email})

	if err != nil {
		return false, err
	}

	if len(users) < 1 {
		return false, nil
	}

	return true, nil
}

func (srv *IUserSerivce) SaveUser(user models.Account) (models.Account, error) {
	return srv.InsertUser(user)
}

func (srv *IUserSerivce) CheckPermission(role_id uint, permissionName, scope string) (bool, error) {
	permission, err := srv.FindPermission(models.Permission{Name: permissionName, Scope: scope})

	if err != nil {
		return false, err
	}

	if len(permission) < 1 {
		return false, nil
	}

	grant, err := srv.FindGrant(models.Rolepermission{Role_id: role_id, Permission_id: permission[0].Permission_id})

	if err != nil {
		return false, err
	}

	if len(grant) < 1 {
		return false, nil
	}

	return true, nil
}

func (srv *IUserSerivce) CheckRole(role_id uint, role string) (bool, error) {
	roles, err := srv.FindRole(models.Role{Role_id: role_id})

	if err != nil {
		return false, err
	}

	if len(roles) < 1 {
		return false, nil
	}

	if roles[0].Name != role {
		return false, nil
	}

	return true, nil
}

func (srv *IUserSerivce) GetUserByID(id uint) (models.Account, error) {
	users, err := srv.FindUser(models.Account{User_id: id})

	if err != nil {
		return models.Account{}, err
	}

	if len(users) < 1 {
		return models.Account{}, errors.New("no users founded")
	}

	return users[0], nil
}

func (srv *IUserSerivce) GetUserByEmail(email string) (models.Account, error) {
	users, err := srv.FindUser(models.Account{Email: email})

	if err != nil {
		return models.Account{}, err
	}

	if len(users) < 1 {
		return models.Account{}, errors.New("no users founded")
	}

	return users[0], nil
}

func (srv *IUserSerivce) GetAllUser() ([]models.Account, error) {
	users, err := srv.FindUser(models.Account{})

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (srv *IUserSerivce) DeleteUserByID(id uint) error {
	err := srv.DeleteUser(models.Account{User_id: id})

	if err != nil {
		return err
	}

	return nil
}

func (srv *IUserSerivce) UpdateUserByID(user models.Account) (models.Account, error) {
	user, err := srv.UpdateUser(user)

	if err != nil {
		return models.Account{}, err
	}

	return user, nil
}
