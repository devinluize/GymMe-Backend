package EquipmentServiceImpl

import (
	entities "GymMe-Backend/api/entities/Equipment"
	"GymMe-Backend/api/helper"
	"GymMe-Backend/api/payloads/responses"
	menuRepository "GymMe-Backend/api/repositories/Equipment"
	"GymMe-Backend/api/service/EquipmentService"
	"gorm.io/gorm"
)

type EquipmentBookmarkServiceImpl struct {
	repository menuRepository.EquipmentBookmarkRepository
	db         *gorm.DB
}

func (e *EquipmentBookmarkServiceImpl) AddEquipmentBookmark(userId, equipmentCourseId int) (entities.EquipmentBookmark, *responses.ErrorResponses) {
	trans := e.db.Begin()
	res, err := e.repository.AddEquipmentBookmark(trans, userId, equipmentCourseId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (e *EquipmentBookmarkServiceImpl) RemoveEquipmentBookmark(userId, equipmentCourseId int) (bool, *responses.ErrorResponses) {
	trans := e.db.Begin()
	res, err := e.repository.RemoveEquipmentBookmark(trans, userId, equipmentCourseId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func NewEquipmentBookmarkServiceImpl(repository menuRepository.EquipmentBookmarkRepository, db *gorm.DB) EquipmentService.EquipmentBookmarkService {
	return &EquipmentBookmarkServiceImpl{repository: repository, db: db}

}
