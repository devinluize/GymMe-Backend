package MenuPayloads

import "time"

type InformationBodyDetail struct {
	InformationBodyParagraph    string `gorm:"column:information_body_paragraph" json:"information_body_paragraph"`
	InformationImageContentPath string `gorm:"column:information_image_content_path" json:"information_image_content_path"`
}
type InformationInsertPayloads struct {
	InformationId                int                     `json:"information_id"`
	InformationHeader            string                  `json:"information_header"`
	InformationImageContentPath1 string                  `json:"information_image_content_path_1"`
	InformationImageContentPath2 string                  `json:"information_image_content_path_2"`
	InformationImageContentPath3 string                  `json:"information_image_content_path_3"`
	InformationImageContentPath4 string                  `json:"information_image_content_path_4"`
	InformationImageContentPath5 string                  `json:"information_image_content_path_5"`
	InformationBodyParagraph     []InformationBodyDetail `json:"information_body_paragraph"`
	InformationTypeId            int                     `json:"information_type_id"`
}
type InformationSelectResponses struct {
	InformationHeader          string                  `json:"information_header"`
	InformationDateCreated     time.Time               `gorm:"column:information_date_created" json:"information_date_created"`
	InformationCreatedByUserId int                     `gorm:"column:information_created_by_user_id" json:"information_created_by_user_id"`
	InformationBodyContent     []InformationBodyDetail `json:"information_body_content"`
	InformationTypeId          int                     `json:"information_type_id"`
}
type InformationUpdatePayloads struct {
	InformationId          int                     `json:"information_id"`
	InformationBodyContent []InformationBodyDetail `json:"information_body_content"`
}
