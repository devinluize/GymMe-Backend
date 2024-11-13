package menuserviceimpl

import (
	"GymMe-Backend/api/entities"
	"GymMe-Backend/api/helper"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	menuRepository "GymMe-Backend/api/repositories/menu"
	"GymMe-Backend/api/service/menu"
	"gorm.io/gorm"
)

type TimerServiceImpl struct {
	repository menuRepository.TimerRepository
	DB         *gorm.DB
}

func NewTimerServiceImpl(repository menuRepository.TimerRepository, db *gorm.DB) menu.TimerService {
	return &TimerServiceImpl{DB: db, repository: repository}
}

func (t *TimerServiceImpl) InsertTimer(payload MenuPayloads.TimerInsertResponse) (entities.TimerEntity, *responses.ErrorResponses) {
	db := t.DB.Begin()
	defer helper.CommitOrRollback(db)
	res, err := t.repository.InsertTimer(db, payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (t *TimerServiceImpl) InsertQueueTimer(payload MenuPayloads.TimerQueueInsertResponse) (entities.TimerQueueEntity, *responses.ErrorResponses) {
	db := t.DB.Begin()
	defer helper.CommitOrRollback(db)
	res, err := t.repository.InsertQueueTimer(db, payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (t *TimerServiceImpl) UpdateQueueTimer(payload MenuPayloads.TimerQueueUpdatePayload) (entities.TimerQueueEntity, *responses.ErrorResponses) {
	db := t.DB.Begin()
	defer helper.CommitOrRollback(db)
	res, err := t.repository.UpdateQueueTimer(db, payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (t *TimerServiceImpl) DeleteTimerQueueTimer(TimerQueueId int) (bool, *responses.ErrorResponses) {
	db := t.DB.Begin()
	defer helper.CommitOrRollback(db)
	res, err := t.repository.DeleteTimerQueueTimer(db, TimerQueueId)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (t *TimerServiceImpl) GetAllTimer(timerId int) ([]entities.TimerEntity, *responses.ErrorResponses) {
	db := t.DB.Begin()
	defer helper.CommitOrRollback(db)
	res, err := t.repository.GetAllTimer(db, timerId)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (t *TimerServiceImpl) GetAllQueueTimer(timerQueueId int) ([]entities.TimerQueueEntity, *responses.ErrorResponses) {
	db := t.DB.Begin()
	defer helper.CommitOrRollback(db)
	res, err := t.repository.GetAllQueueTimer(db, timerQueueId)
	if err != nil {
		return res, err
	}
	return res, nil
}
