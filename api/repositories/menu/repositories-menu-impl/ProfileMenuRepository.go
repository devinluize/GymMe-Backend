package MenuImplRepositories

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	menuRepository "GymMe-Backend/api/repositories/menu"
	"gorm.io/gorm"
	"net/http"
)

type ProfileMenuRepositoryImpl struct {
}

func NewProfileMenuRepositoryImpl() menuRepository.ProfileMenuRepository {
	return &ProfileMenuRepositoryImpl{}

}
func (p *ProfileMenuRepositoryImpl) GetProfileMenu(db *gorm.DB, id int) (entities.UserDetail, *responses.ErrorResponses) {
	Entities := entities.UserDetail{}
	err := db.Where(entities.UserDetail{UserId: id}).First(&Entities).Error
	if err != nil {
		return Entities,
			&responses.ErrorResponses{StatusCode: http.StatusInternalServerError,
				Err:     err,
				Message: err.Error()}
	}
	return Entities, nil
}

func (p *ProfileMenuRepositoryImpl) UpdateProfileMenu(db *gorm.DB, Request MenuPayloads.ProfilePayloadRequest) (entities.UserDetail, *responses.ErrorResponses) {
	Entities := entities.UserDetail{}
	err := db.Model(&Entities).Where(entities.UserDetail{UserId: Request.UserId}).Error
	if err != nil {
		return Entities, &responses.ErrorResponses{StatusCode: http.StatusInternalServerError,
			Err:     err,
			Message: err.Error()}
	}
	Entities.UserHeight = Request.UserHeight
	Entities.UserWeight = Request.UserWeight
	Entities.UserGender = Request.UserGender

	err = db.Save(&Entities).Error
	if err != nil {
		return Entities, &responses.ErrorResponses{StatusCode: http.StatusInternalServerError,
			Message: "Error updating profile menu",
			Err:     err,
		}
	}
	return Entities, nil
}
func (p *ProfileMenuRepositoryImpl) CreateProfileMenu(db *gorm.DB, Request MenuPayloads.ProfilePayloadRequest) (entities.UserDetail, *responses.ErrorResponses) {
	Entities := entities.UserDetail{
		UserId:                 Request.UserId,
		UserWeight:             Request.UserWeight,
		UserHeight:             Request.UserHeight,
		UserGender:             Request.UserGender,
		UserProfileDescription: Request.UserProfileDescription,
		UserProfileImage:       Request.UserProfileImage,
	}
	err := db.Create(&Entities).Error
	if err != nil {
		return Entities, &responses.ErrorResponses{StatusCode: http.StatusBadRequest,
			Message: "Error creating profile menu",
			Err:     err,
		}
	}
	return Entities, nil
}
