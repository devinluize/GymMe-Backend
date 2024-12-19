package EquipmentServiceImpl

import (
	entities "GymMe-Backend/api/entities/Equipment"
	"GymMe-Backend/api/helper"
	"GymMe-Backend/api/payloads/Equipment"
	"GymMe-Backend/api/payloads/responses"
	menuRepository "GymMe-Backend/api/repositories/Equipment"
	"GymMe-Backend/api/service/EquipmentService"
	"gorm.io/gorm"
)

type EquipmentCourseServiceImpl struct {
	repository menuRepository.EquipmentCourseRepository
	db         *gorm.DB
}

func NewEquipmentCourseServiceImpl(db *gorm.DB, repository menuRepository.EquipmentCourseRepository) EquipmentService.EquipmentCourseService {
	return &EquipmentCourseServiceImpl{
		db:         db,
		repository: repository,
	}
}
func (s *EquipmentCourseServiceImpl) GetAllEquipmentCourseByEquipment(equipmentId int) (Equipment.GetAllCourseEquipmentResponse, *responses.ErrorResponses) {
	trans := s.db.Begin()
	res, err := s.repository.GetAllEquipmentCourseByEquipment(trans, equipmentId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *EquipmentCourseServiceImpl) InsertEquipmentCourse(payload Equipment.InsertEquipmentCourseDataPayload) (entities.EquipmentCourseDataEntity, *responses.ErrorResponses) {
	trans := s.db.Begin()
	res, err := s.repository.InsertEquipmentCourse(trans, payload)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *EquipmentCourseServiceImpl) GetEquipmentCourse(courseId int) (Equipment.GetCourseByIdResponse, *responses.ErrorResponses) {
	trans := s.db.Begin()
	res, err := s.repository.GetEquipmentCourse(trans, courseId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
