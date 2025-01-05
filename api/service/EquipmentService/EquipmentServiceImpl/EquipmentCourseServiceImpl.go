package EquipmentServiceImpl

import (
	entities "GymMe-Backend/api/entities/Equipment"
	"GymMe-Backend/api/helper"
	"GymMe-Backend/api/payloads/Equipment"
	"GymMe-Backend/api/payloads/responses"
	menuRepository "GymMe-Backend/api/repositories/Equipment"
	"GymMe-Backend/api/service/EquipmentService"
	"github.com/cloudinary/cloudinary-go/v2"
	"gorm.io/gorm"
)

type EquipmentCourseServiceImpl struct {
	repository menuRepository.EquipmentCourseRepository
	db         *gorm.DB
	cld        *cloudinary.Cloudinary
}

func NewEquipmentCourseServiceImpl(db *gorm.DB, repository menuRepository.EquipmentCourseRepository, cld *cloudinary.Cloudinary) EquipmentService.EquipmentCourseService {
	return &EquipmentCourseServiceImpl{
		db:         db,
		repository: repository,
		cld:        cld,
	}
}
func (s *EquipmentCourseServiceImpl) GetAllEquipmentCourseByEquipment(equipmentId int) (Equipment.GetAllCourseEquipmentResponse, *responses.ErrorResponses) {
	trans := s.db.Begin()
	res, err := s.repository.GetAllEquipmentCourseByEquipment(trans, equipmentId, s.cld)
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
	res, err := s.repository.GetEquipmentCourse(trans, courseId, s.cld)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *EquipmentCourseServiceImpl) SearchEquipmentByKey(EquipmentKey string, userId int) ([]entities.EquipmentMasterEntities, *responses.ErrorResponses) {
	trans := s.db.Begin()
	res, err := s.repository.SearchEquipmentByKey(trans, EquipmentKey, userId, s.cld)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *EquipmentCourseServiceImpl) GetEquipmentSearchHistoryByKey(userId int) ([]entities.EquipmentSearchHistoryEntities, *responses.ErrorResponses) {
	trans := s.db.Begin()
	res, err := s.repository.GetEquipmentSearchHistoryByKey(trans, userId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *EquipmentCourseServiceImpl) DeleteEquipmentSearchHistoryById(equipmentSearchHistoryId int) (bool, *responses.ErrorResponses) {
	trans := s.db.Begin()
	res, err := s.repository.DeleteEquipmentSearchHistoryById(trans, equipmentSearchHistoryId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
