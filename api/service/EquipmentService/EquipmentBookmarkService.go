package EquipmentService

import (
	entities "GymMe-Backend/api/entities/Equipment"
	"GymMe-Backend/api/payloads/Equipment"
	"GymMe-Backend/api/payloads/responses"
)

type EquipmentBookmarkService interface {
	AddEquipmentBookmark(userId, equipmentCourseId int) (entities.EquipmentBookmark, *responses.ErrorResponses)
	RemoveEquipmentBookmark(userId, equipmentCourseId int) (bool, *responses.ErrorResponses)
	GetEquipmentBookmarkByUserId(userId int) ([]Equipment.GetBookmarkEquipmentResponse, *responses.ErrorResponses)
}
