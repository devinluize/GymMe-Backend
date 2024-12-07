package entities

type EquipmentDetailEntity struct {
	EquipmentDetailId          int    `gorm:"column:equipment_detail_id;primaryKey"`
	EquipmentId                int    `gorm:"column:equipment_id"`
	EquipmentDetailHeader      string `gorm:"column:equipment_detail_header"`
	EquipmentDetailDescription string `gorm:"column:equipment_detail_description"`
}
