package menuRepository

import (
	"GymMe-Backend/api/entities"
	"GymMe-Backend/api/helper"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	"github.com/cloudinary/cloudinary-go/v2"
	"gorm.io/gorm"
)

type InformationMenu interface {
	InsertInformation(tx *gorm.DB, payloads MenuPayloads.InformationInsertPayloads) (entities.InformationEntities, *responses.ErrorResponses)
	DeleteInformationById(db *gorm.DB, id int) (bool, *responses.ErrorResponses)
	UpdateInformation(tx *gorm.DB, payloads MenuPayloads.InformationUpdatePayloads) (entities.InformationEntities, *responses.ErrorResponses)
	GetAllInformationWithPagination(db *gorm.DB, paginationResponses helper.Pagination) (helper.Pagination, *responses.ErrorResponses)
	GetInformationById(db *gorm.DB, id int, userId int, cld *cloudinary.Cloudinary) (MenuPayloads.InformationSelectResponses, *responses.ErrorResponses)
	GetAllInformationWithFilter(db *gorm.DB, paginationResponses helper.Pagination, Key string, userId int, cloudinary *cloudinary.Cloudinary) (helper.Pagination, *responses.ErrorResponses)
	GetInformationHistory(db *gorm.DB, userId int) ([]entities.SearchHistoryEntities, *responses.ErrorResponses)
}
