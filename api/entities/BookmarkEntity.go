package entities

type Bookmark struct {
	BookmarkId     int `gorm:"column:bookmark_id;not null;primaryKey" json:"bookmark_id"`
	BookmarkTypeId int `gorm:"column:bookmark_type_id" json:"bookmark_type_id"`
	BookmarkType   BookmarkType
	InformationId  int `gorm:"column:information_id" json:"information_id"`
	Information    InformationEntities
	UserId         int `gorm:"column:user_id" json:"user_id"`
	User           Users
}

func (*Bookmark) TableName() string {
	return "mtr_bookmark"
}
