package EquipmentService

import (
	entities "GymMe-Backend/api/entities/Equipment"
	"GymMe-Backend/api/payloads/Equipment"
	"GymMe-Backend/api/payloads/responses"
)

type EquipmentCourseService interface {
	GetAllEquipmentCourseByEquipment(equipmentId int) (Equipment.GetAllCourseEquipmentResponse, *responses.ErrorResponses)
	InsertEquipmentCourse(payload Equipment.InsertEquipmentCourseDataPayload) (entities.EquipmentCourseDataEntity, *responses.ErrorResponses)
	GetEquipmentCourse(courseId int) (Equipment.GetCourseByIdResponse, *responses.ErrorResponses)
}
