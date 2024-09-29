package entities

import "time"

type InformationEntities struct {
	InformationId              int                       `gorm:"column:information_id;not null;primaryKey" json:"information_id"`
	InformationHeader          string                    `gorm:"column:information_header" json:"information_header"`
	InformationDateCreated     time.Time                 `gorm:"column:information_date_created" json:"information_date_created"`
	InformationCreatedByUserId int                       `gorm:"column:information_created_by_user_id" json:"information_created_by_user_id"`
	InformationBody            []InformationBodyEntities `gorm:"foreignKey:InformationId;references:InformationId" json:"information_body"`
	InformationTypeId          int                       `gorm:"column:information_type_id" json:"information_type_id"`
	InformationType            InformationType
}

func (*InformationEntities) TableName() string {
	return "mtr_information"
}
