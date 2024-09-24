package entities

type InformationType struct {
	InformationTypeId   int    `gorm:"column:information_type_id;primaryKey;not null" json:"information_type_id"`
	InformationTypeName string `gorm:"information_type_name" json:"information_type_name"`
}

func (*InformationType) TableName() string {
	return "mtr_information_type"
}
