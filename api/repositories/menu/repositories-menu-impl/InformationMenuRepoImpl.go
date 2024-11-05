package MenuImplRepositories

import (
	"GymMe-Backend/api/entities"
	"GymMe-Backend/api/helper"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	menuRepository "GymMe-Backend/api/repositories/menu"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

type InformationMenu struct {
}

func NewInformationMenu() menuRepository.InformationMenu {
	return &InformationMenu{}
}
func (i *InformationMenu) UpdateInformation(tx *gorm.DB, payloads MenuPayloads.InformationUpdatePayloads) (entities.InformationEntities, *responses.ErrorResponses) {
	var InformationEntities entities.InformationEntities
	var InformationBodyEntities []entities.InformationBodyEntities

	err := tx.Model(&InformationEntities).Where(entities.InformationEntities{InformationId: payloads.InformationId}).First(&InformationEntities).Error
	if err != nil {
		return InformationEntities,
			&responses.ErrorResponses{StatusCode: http.StatusInternalServerError,
				Err:     err,
				Message: err.Error()}
	}
	for _, i := range payloads.InformationBodyContent {
		InformationBodyEntitiesData := entities.InformationBodyEntities{
			InformationBodyParagraph:    i.InformationBodyParagraph,
			InformationImageContentPath: i.InformationImageContentPath,
			InformationId:               payloads.InformationId,
		}
		InformationBodyEntities = append(InformationBodyEntities, InformationBodyEntitiesData)
	}
	//Information Body Inserting
	InformationEntities.InformationBody = InformationBodyEntities
	err = tx.Delete(entities.InformationBodyEntities{}, entities.InformationBodyEntities{InformationId: payloads.InformationId}).Error
	err = tx.Save(&InformationEntities).Error
	if err != nil {
		return InformationEntities, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			//Err:        nil,
			Success: false,
		}
	}
	return InformationEntities, nil
}
func (i *InformationMenu) DeleteInformationById(db *gorm.DB, id int) (bool, *responses.ErrorResponses) {
	err := db.Model(&entities.InformationEntities{}).Where(entities.InformationEntities{InformationId: id}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, &responses.ErrorResponses{
				StatusCode: http.StatusBadRequest,
				Message:    "Delete Failed Id Not Found",
				Success:    false,
			}
		}
		return false, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	err = db.Delete(entities.InformationBodyEntities{}, entities.InformationBodyEntities{InformationId: id}).Error
	if err != nil {
		return false, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}

	err = db.Delete(entities.InformationEntities{}, entities.InformationEntities{InformationId: id}).Error
	if err != nil {
		return false, &responses.ErrorResponses{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}
	return true, nil
}
func (i *InformationMenu) InsertInformation(tx *gorm.DB, payloads MenuPayloads.InformationInsertPayloads) (entities.InformationEntities, *responses.ErrorResponses) {
	var Entities entities.InformationEntities
	Entities = entities.InformationEntities{
		//InformationId:                0,
		InformationHeader: payloads.InformationHeader,
		//InformationImageContentPath1: payloads.InformationImageContentPath1,
		//InformationImageContentPath2: payloads.InformationImageContentPath2,
		//InformationImageContentPath3: payloads.InformationImageContentPath3,
		//InformationImageContentPath4: payloads.InformationImageContentPath4,
		//InformationImageContentPath5: payloads.InformationImageContentPath5,
		InformationTypeId: payloads.InformationTypeId,
	}
	var EntitiesDetail []entities.InformationBodyEntities

	//Entities.InformationId = Entities.InformationId
	err := tx.Create(&Entities).Error
	if err != nil {
		return Entities, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Err:        err,
			Success:    false,
		}
	}

	err = tx.Model(&Entities).Where(entities.InformationEntities{InformationId: Entities.InformationId}).First(&Entities).Error
	for _, paragraph := range payloads.InformationBodyParagraph {
		EntitiesDetail = append(EntitiesDetail, entities.InformationBodyEntities{InformationBodyParagraph: paragraph.InformationBodyParagraph,
			InformationId: Entities.InformationId,
		})
	}
	err = tx.Create(&EntitiesDetail).Error
	if err != nil {
		return Entities, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Err:        err,
			Success:    false,
		}
	}
	return Entities, nil
}
func (i *InformationMenu) GetInformationById(db *gorm.DB, id int) (MenuPayloads.InformationSelectResponses, *responses.ErrorResponses) {
	EntitiesInfo := entities.InformationEntities{}
	err := db.Model(&entities.InformationEntities{}).Where(entities.InformationEntities{InformationId: id}).First(&EntitiesInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return MenuPayloads.InformationSelectResponses{}, &responses.ErrorResponses{
				StatusCode: http.StatusBadRequest,
				Message:    "Delete Failed Id Not Found",
			}
		}
	}
	ResDetail := []MenuPayloads.InformationBodyDetail{}
	err = db.Model(&entities.InformationBodyEntities{}).Where(entities.InformationBodyEntities{InformationId: id}).Scan(&ResDetail).Error
	if err != nil {
		return MenuPayloads.InformationSelectResponses{}, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}
	SelectPayload := []MenuPayloads.InformationBodyDetail{}
	for _, detail := range ResDetail {
		SelectPayloadData := MenuPayloads.InformationBodyDetail{
			InformationBodyParagraph:    detail.InformationBodyParagraph,
			InformationImageContentPath: detail.InformationImageContentPath,
		}
		SelectPayload = append(SelectPayload, SelectPayloadData)
	}

	result := MenuPayloads.InformationSelectResponses{
		InformationHeader:          EntitiesInfo.InformationHeader,
		InformationDateCreated:     EntitiesInfo.InformationDateCreated,
		InformationCreatedByUserId: EntitiesInfo.InformationCreatedByUserId,
		InformationBodyContent:     SelectPayload,
		InformationTypeId:          EntitiesInfo.InformationTypeId,
	}
	return result, nil
}
func (i *InformationMenu) GetAllInformationWithPagination(db *gorm.DB, paginationResponses helper.Pagination) (helper.Pagination, *responses.ErrorResponses) {

	var Entities []entities.InformationEntities
	//me := db.Model(&entities.InformationEntities{}) -> table joinan
	//cara 1
	//err := db.Model(&entities.InformationEntities{}).Scopes(database.Paginate(&Entities, &paginationResponses, me)).Order("information_id").Where("information_id <> 0").Scan(&Entities).Error
	//cara 2 langsung assign ke database nanti pilih aja apakah perlu buat join table atau ga kalau misalkan selectan itu merupakan hasil join table pake yang atas
	err := db.Model(&entities.InformationEntities{}).Scopes(helper.Paginate(&Entities, &paginationResponses, db)).Order("information_id").Where("information_id <> 0").Scan(&Entities).Error
	if err != nil {
		return paginationResponses, &responses.ErrorResponses{}
	}
	paginationResponses.Rows = Entities
	fmt.Println(paginationResponses.Rows)
	return paginationResponses, nil
}
func (i *InformationMenu) GetAllInformationWithFilter(db *gorm.DB, paginationResponses helper.Pagination, Key string) (helper.Pagination, *responses.ErrorResponses) {
	var Entities []entities.InformationEntities
	//me := db.Model(&entities.InformationEntities{}) -> table joinan
	//cara 1
	//err := db.Model(&entities.InformationEntities{}).Scopes(database.Paginate(&Entities, &paginationResponses, me)).Order("information_id").Where("information_id <> 0").Scan(&Entities).Error
	//cara 2 langsung assign ke database nanti pilih aja apakah perlu buat join table atau ga kalau misalkan selectan itu merupakan hasil join table pake yang atas
	err := db.Model(&entities.InformationEntities{}).Scopes(helper.Paginate(&Entities, &paginationResponses, db)).Order("information_id").Where("information_id <> 0 AND information_header LIKE ? ", "%"+Key+"%").Scan(&Entities).Error
	if err != nil {
		return paginationResponses, &responses.ErrorResponses{}
	}
	paginationResponses.Rows = Entities
	fmt.Println(paginationResponses.Rows)
	return paginationResponses, nil
}
