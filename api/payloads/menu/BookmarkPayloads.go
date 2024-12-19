package MenuPayloads

import "time"

type GetAllBookmarkResponse struct {
	InformationHeader            string    `json:"information_header"`
	InformationDateCreated       time.Time `gorm:"column:information_date_created" json:"information_date_created"`
	InformationCreatedByUserId   int       `gorm:"column:information_created_by_user_id" json:"information_created_by_user_id"`
	InformationId                int       `gorm:"column:information_id;not null;primaryKey" json:"information_id"`
	BookmarkId                   int       `json:"bookmark_id"`
	InformationHeaderPathContent string    `gorm:"column:information_header_path_content" json:"information_header_path_content"`
}
