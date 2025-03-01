package menuRepository

import (
	entities "GymMe-Backend/api/entities/Equipment"
	"GymMe-Backend/api/payloads/Equipment"
	"GymMe-Backend/api/payloads/responses"
	"github.com/cloudinary/cloudinary-go/v2"
	"gorm.io/gorm"
)

type EquipmentBookmarkRepository interface {
	AddEquipmentBookmark(db *gorm.DB, userId, equipmentCourseId int) (entities.EquipmentBookmark, *responses.ErrorResponses)
	RemoveEquipmentBookmark(db *gorm.DB, userId, equipmentCourseId int) (bool, *responses.ErrorResponses)
	GetEquipmentBookmarkByUserId(db *gorm.DB, userId int, cld *cloudinary.Cloudinary) ([]Equipment.GetBookmarkEquipmentResponse, *responses.ErrorResponses)
}
