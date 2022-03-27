package services

import (
	"challenge4/models"
	"challenge4/repositories"
)

type IGrantService struct {
	repositories.GrantRepository
}

func (srv *IGrantService) GetAllGrant() ([]models.Rolepermission, error) {
	grants, err := srv.FindGrant(models.Rolepermission{})

	if err != nil {
		return nil, err
	}

	return grants, nil
}
func (srv *IGrantService) DeleteGrantSrv(grant models.Rolepermission) error {
	return srv.DeleteGrant(grant)
}
func (srv *IGrantService) InsertGrant(grant models.Rolepermission) (models.Rolepermission, error) {
	return srv.SaveGrant(grant)
}
