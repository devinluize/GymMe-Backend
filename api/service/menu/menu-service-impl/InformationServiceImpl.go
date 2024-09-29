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

type InformationServiceImpl struct {
	repo menuRepository.InformationMenu
	db   *gorm.DB
}

func NewInformationServiceImpl(repo menuRepository.InformationMenu, db *gorm.DB) menu.InformationService {
	return &InformationServiceImpl{repo: repo, db: db}
}
func (service *InformationServiceImpl) InsertInformation(payloads MenuPayloads.InformationInsertPayloads) (entities.InformationEntities, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repo.InsertInformation(trans, payloads)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *InformationServiceImpl) DeleteInformationById(id int) (bool, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repo.DeleteInformationById(trans, id)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *InformationServiceImpl) UpdateInformation(payloads MenuPayloads.InformationUpdatePayloads) (entities.InformationEntities, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repo.UpdateInformation(trans, payloads)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *InformationServiceImpl) GetInformationById(id int) (MenuPayloads.InformationSelectResponses, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repo.GetInformationById(trans, id)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
