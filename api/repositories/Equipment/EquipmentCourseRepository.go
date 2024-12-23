package menuRepository

import (
	entities "GymMe-Backend/api/entities/Equipment"
	"GymMe-Backend/api/payloads/Equipment"
	"GymMe-Backend/api/payloads/responses"
	"gorm.io/gorm"
)

type EquipmentCourseRepository interface {
	GetAllEquipmentCourseByEquipment(db *gorm.DB, equipmentId int) (Equipment.GetAllCourseEquipmentResponse, *responses.ErrorResponses)
	InsertEquipmentCourse(db *gorm.DB, payload Equipment.InsertEquipmentCourseDataPayload) (entities.EquipmentCourseDataEntity, *responses.ErrorResponses)
	GetEquipmentCourse(db *gorm.DB, courseId int) (Equipment.GetCourseByIdResponse, *responses.ErrorResponses)
	SearchEquipmentByKey(db *gorm.DB, EquipmentKey string) ([]entities.EquipmentMasterEntities, *responses.ErrorResponses)
}
