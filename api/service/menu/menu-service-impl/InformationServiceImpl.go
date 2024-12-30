package menuserviceimpl

import (
	"GymMe-Backend/api/entities"
	"GymMe-Backend/api/helper"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	menuRepository "GymMe-Backend/api/repositories/menu"
	"GymMe-Backend/api/service/menu"
	"github.com/cloudinary/cloudinary-go/v2"
	"gorm.io/gorm"
)

type InformationServiceImpl struct {
	repo menuRepository.InformationMenu
	db   *gorm.DB
	cld  *cloudinary.Cloudinary
}

func NewInformationServiceImpl(repo menuRepository.InformationMenu, db *gorm.DB, cld *cloudinary.Cloudinary) menu.InformationService {
	return &InformationServiceImpl{repo: repo, db: db, cld: cld}
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

func (service *InformationServiceImpl) GetInformationById(id int, userId int) (MenuPayloads.InformationSelectResponses, *responses.ErrorResponses) {
	trans := service.db.Begin()

	res, err := service.repo.GetInformationById(trans, id, userId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (service *InformationServiceImpl) GetAllInformationWithPagination(paginationResponses helper.Pagination) (helper.Pagination, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repo.GetAllInformationWithPagination(trans, paginationResponses)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (service *InformationServiceImpl) GetAllInformationWithFilter(paginationResponses helper.Pagination, Key string, userId int) (helper.Pagination, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repo.GetAllInformationWithFilter(trans, paginationResponses, Key, userId, service.cld)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (service *InformationServiceImpl) GetInformationHistory(userId int) ([]entities.SearchHistoryEntities, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repo.GetInformationHistory(trans, userId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
