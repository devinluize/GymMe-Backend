package entities

const mucleGroupTableName = "mtr_muscle_group"

type MuscleGroupEntities struct {
	MuscleGroupId   int    `gorm:"column:muscle_group_id;not null;primaryKey"`
	MuscleGroupName string `gorm:"column:muscle_group_name;not null"`
}

func (*MuscleGroupEntities) TableName() string {
	return mucleGroupTableName
}
