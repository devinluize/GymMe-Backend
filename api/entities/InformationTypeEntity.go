package entities

type InformationType struct {
	InformationTypeEntities int    `gorm:"column:information_type_entities;primaryKey;not null" json:"information_type_entities"`
	InformationTypeName     string `gorm:"information_type_name" json:"information_type_name"`
}

func (*InformationType) TableName() string {
	return "mtr_information_type"
}
