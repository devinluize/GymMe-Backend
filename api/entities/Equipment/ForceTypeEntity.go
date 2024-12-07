package entities

const ForceTypeTableName = "mtr_force_type"

type ForceTypeEntities struct {
	ForceTypeId   int    `gorm:"column:force_type_id;not null;primaryKey"`
	ForceTypeName string `gorm:"column:force_type_name;not null"`
}

func (*ForceTypeEntities) TableName() string {
	return ForceTypeTableName
}
