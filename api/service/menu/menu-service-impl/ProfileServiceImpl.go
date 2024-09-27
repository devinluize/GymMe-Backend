package menuserviceimpl

import (
	"GymMe-Backend/api/entities"
	"GymMe-Backend/api/helper"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	menuRepository "GymMe-Backend/api/repositories/menu"
	"GymMe-Backend/api/service/menu"
	"gorm.io/gorm"
)

func NewProfileServiceImpl(db *gorm.DB, repository menuRepository.ProfileMenuRepository) menu.ProfileService {
	return &ProfileServiceImpl{
		db:         db,
		repository: repository,
	}

}

type ProfileServiceImpl struct {
	db         *gorm.DB
	repository menuRepository.ProfileMenuRepository
}

func (service *ProfileServiceImpl) GetProfileMenu(id int) (entities.UserDetail, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repository.GetProfileMenu(trans, id)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *ProfileServiceImpl) UpdateProfileMenu(Request MenuPayloads.ProfilePayloadRequest) (entities.UserDetail, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repository.UpdateProfileMenu(trans, Request)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *ProfileServiceImpl) CreateProfileMenu(Request MenuPayloads.ProfilePayloadRequest) (entities.UserDetail, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repository.CreateProfileMenu(trans, Request)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
