package entities

type InformationBodyEntities struct {
	InformationBodyId           int    `gorm:"column:information_body_id;primaryKey;not null" json:"information_body_id"`
	InformationId               int    `gorm:"column:information_id" json:"information_id"`
	InformationBodyParagraph    string `gorm:"column:information_body_paragraph" json:"information_body_paragraph"`
	InformationImageContentPath string `gorm:"column:information_image_content_path" json:"information_image_content_path"`
}

func (*InformationBodyEntities) TableName() string {
	return "mtr_information_detail_body"
}
