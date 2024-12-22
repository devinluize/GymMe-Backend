package MenuImplRepositories

import (
	"GymMe-Backend/api/entities"
	payloads "GymMe-Backend/api/payloads/auth"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	menuRepository "GymMe-Backend/api/repositories/menu"
	"errors"
	"gorm.io/gorm"
	"math"
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
	err := db.Model(&Entities).Where(entities.UserDetail{UserId: Request.UserId}).Scan(&Entities).Error
	if err != nil {
		return Entities, &responses.ErrorResponses{StatusCode: http.StatusInternalServerError,
			Err:     err,
			Message: err.Error()}
	}
	Entities.UserHeight = Request.UserHeight
	Entities.UserWeight = Request.UserWeight
	Entities.UserGender = Request.UserGender

	err = db.Updates(&Entities).Error
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
func (p *ProfileMenuRepositoryImpl) GetBmi(db *gorm.DB, userId int) (payloads.UserBmiResponse, *responses.ErrorResponses) {
	//get user weight first with user id
	userBmiResponse := payloads.UserBmiResponse{}
	var lastWeight float64
	err := db.Model(&entities.WeightHistoryEntities{}).
		Where(entities.WeightHistoryEntities{UserId: userId}).
		Select("user_weight").Limit(1).Order("user_weight_time desc").
		Scan(&lastWeight).Error
	if err != nil {
		return userBmiResponse, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to get user weight by user id",
		}
	}
	//get profile
	profile := entities.UserDetail{}
	err = db.Model(&profile).Where(entities.UserDetail{UserId: userId}).First(&profile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userBmiResponse, &responses.ErrorResponses{
				StatusCode: http.StatusNotFound,
				Err:        err,
				Message:    "user detail is not found",
			}
		}
		return userBmiResponse, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to get user detail",
			Err:        err,
		}
	}
	userHeightMeter := profile.UserHeight / 100

	Bmi := lastWeight / (math.Pow(userHeightMeter, 2))
	userBmiResponse.Bmi = Bmi
	userBmiResponse.UserId = userId
	userBmiResponse.UserHeight = userHeightMeter
	userBmiResponse.UserWeight = lastWeight
	return userBmiResponse, nil
}
