package MenuPayloads

import "time"

type InformationBodyDetail struct {
	InformationBodyParagraph    string `gorm:"column:information_body_paragraph" json:"information_body_paragraph"`
	InformationImageContentPath string `gorm:"column:information_image_content_path" json:"information_image_content_path"`
}
type InformationInsertPayloads struct {
	//InformationId                int                     `json:"information_id"`
	InformationHeader            string                  `json:"information_header"`
	InformationHeaderPathContent string                  `gorm:"column:information_header_path_content" json:"information_header_path_content"`
	InformationBodyParagraph     []InformationBodyDetail `json:"information_body_paragraph"`
	//InformationTypeId            int                     `json:"information_type_id"`
}
type InformationSelectResponses struct {
	InformationHeader          string    `json:"information_header"`
	InformationDateCreated     time.Time `gorm:"column:information_date_created" json:"information_date_created"`
	InformationCreatedByUserId int       `gorm:"column:information_created_by_user_id" json:"information_created_by_user_id"`
	InformationId              int       `gorm:"column:information_id;not null;primaryKey" json:"information_id"`

	InformationBodyContent []InformationBodyDetail `json:"information_body_content"`
	IsBookmark             bool                    `json:"is_bookmark"`
	//InformationTypeId          int                     `json:"information_type_id"`
}
type InformationSelectResponseHeader struct {
	InformationHeader          string    `json:"information_header"`
	InformationDateCreated     time.Time `gorm:"column:information_date_created" json:"information_date_created"`
	InformationCreatedByUserId int       `gorm:"column:information_created_by_user_id" json:"information_created_by_user_id"`
	InformationId              int       `gorm:"column:information_id;not null;primaryKey" json:"information_id"`
	BookmarkId                 int       `json:"bookmark_id"`
}

type InformationUpdatePayloads struct {
	InformationId          int                     `json:"information_id"`
	InformationBodyContent []InformationBodyDetail `json:"information_body_content"`
}
