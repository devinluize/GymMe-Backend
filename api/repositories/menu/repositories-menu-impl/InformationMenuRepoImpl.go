package MenuImplRepositories

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	menuRepository "GymMe-Backend/api/repositories/menu"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type InformationMenu struct {
}

func NewInformationMenu() menuRepository.InformationMenu {
	return &InformationMenu{}
}

func (i *InformationMenu) UpdateInformation(tx *gorm.DB, payloads MenuPayloads.InformationUpdatePayloads) (entities.Information, *responses.ErrorResponses) {
	var InformationEntities entities.Information

	err := tx.Model(&InformationEntities).Where(&MenuPayloads.InformationUpdatePayloads{InformationId: payloads.InformationId}).
		First(&InformationEntities).Error

	if err != nil {
		return InformationEntities,
			&responses.ErrorResponses{StatusCode: http.StatusInternalServerError,
				Err:     err,
				Message: err.Error()}
	}

	InformationEntities.InformationBodyParagraph1 = payloads.InformationBodyParagraph1
	InformationEntities.InformationBodyParagraph2 = payloads.InformationBodyParagraph2
	InformationEntities.InformationBodyParagraph3 = payloads.InformationBodyParagraph3
	InformationEntities.InformationBodyParagraph4 = payloads.InformationBodyParagraph4
	InformationEntities.InformationImageContentPath1 = payloads.InformationImageContentPath1
	InformationEntities.InformationImageContentPath2 = payloads.InformationImageContentPath2
	InformationEntities.InformationImageContentPath3 = payloads.InformationImageContentPath3
	InformationEntities.InformationImageContentPath4 = payloads.InformationImageContentPath4
	InformationEntities.InformationImageContentPath5 = payloads.InformationImageContentPath5

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
	err := db.Model(&entities.Information{}).Where(entities.Information{InformationId: id}).Error
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
	err = db.Delete(entities.Information{}, entities.Information{InformationId: id}).Error
	if err != nil {
		return false, &responses.ErrorResponses{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}
	return true, nil
}

func (i *InformationMenu) InsertInformation(tx *gorm.DB, payloads MenuPayloads.InformationInsertPayloads) (entities.Information, *responses.ErrorResponses) {
	var Entities entities.Information
	Entities = entities.Information{
		//InformationId:                0,
		InformationHeader:            payloads.InformationHeader,
		InformationImageContentPath1: payloads.InformationImageContentPath1,
		InformationImageContentPath2: payloads.InformationImageContentPath2,
		InformationImageContentPath3: payloads.InformationImageContentPath3,
		InformationImageContentPath4: payloads.InformationImageContentPath4,
		InformationImageContentPath5: payloads.InformationImageContentPath5,
		InformationBodyParagraph1:    payloads.InformationBodyParagraph1,
		InformationBodyParagraph2:    payloads.InformationBodyParagraph2,
		InformationBodyParagraph3:    payloads.InformationBodyParagraph3,
		InformationBodyParagraph4:    payloads.InformationBodyParagraph4,
		InformationTypeId:            payloads.InformationId,
	}

	err := tx.Create(&Entities).Error
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
func (i *InformationMenu) GetInformationById(db *gorm.DB, id int) (MenuPayloads.InformationSelectPayloads, *responses.ErrorResponses) {
	EntitiesInfo := entities.Information{}
	err := db.Model(&entities.Information{}).Where(entities.Information{InformationId: id}).First(&EntitiesInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return MenuPayloads.InformationSelectPayloads{}, &responses.ErrorResponses{
				StatusCode: http.StatusBadRequest,
				Message:    "Delete Failed Id Not Found",
			}
		}
	}
	result := MenuPayloads.InformationSelectPayloads{
		InformationHeader:            EntitiesInfo.InformationHeader,
		InformationImageContentPath1: EntitiesInfo.InformationImageContentPath1,
		InformationImageContentPath2: EntitiesInfo.InformationImageContentPath2,
		InformationImageContentPath3: EntitiesInfo.InformationImageContentPath3,
		InformationImageContentPath4: EntitiesInfo.InformationImageContentPath4,
		InformationImageContentPath5: EntitiesInfo.InformationImageContentPath5,
		InformationBodyParagraph1:    EntitiesInfo.InformationBodyParagraph1,
		InformationBodyParagraph2:    EntitiesInfo.InformationBodyParagraph2,
		InformationBodyParagraph3:    EntitiesInfo.InformationBodyParagraph3,
		InformationBodyParagraph4:    EntitiesInfo.InformationBodyParagraph4,
		InformationTypeId:            EntitiesInfo.InformationTypeId,
	}
	return result, nil
}
func (i *InformationMenu) GetAllInformationWithPagination(db *gorm.DB) ([]entities.Information, *responses.ErrorResponses) {
	//TODO implement me
	panic("implement me")
}
