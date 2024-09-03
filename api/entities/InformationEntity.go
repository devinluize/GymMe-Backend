package entities

type Information struct {
	InformationId                int    `gorm:"column:information_id;not null;primaryKey" json:"information_id"`
	InformationHeader            string `gorm:"column:information_header" json:"information_header"`
	InformationImageContentPath1 string `gorm:"column:information_image_content_path_1" json:"information_image_content_path_1"`
	InformationImageContentPath2 string `gorm:"column:information_image_content_path_2" json:"information_image_content_path_2"`
	InformationImageContentPath3 string `gorm:"column:information_image_content_path_3" json:"information_image_content_path_3"`
	InformationImageContentPath4 string `gorm:"column:information_image_content_path_4" json:"information_image_content_path_4"`
	InformationImageContentPath5 string `gorm:"column:information_image_content_path_5" json:"information_image_content_path_5"`
	InformationBodyParagraph1    string `gorm:"column:information_body_paragraph_1" json:"information_body_paragraph_1"`
	InformationBodyParagraph2    string `gorm:"column:information_body_paragraph_2" json:"information_body_paragraph_2"`
	InformationBodyParagraph3    string `gorm:"column:information_body_paragraph_3" json:"information_body_paragraph_3"`
	InformationTypeId            int    `gorm:"column:information_type_id" json:"information_type_id"`
	InformationType              InformationType
}

func (*Information) TableName() string {
	return "information"
}
