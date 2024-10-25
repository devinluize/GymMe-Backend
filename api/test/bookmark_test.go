package test

import (
	configenv "GymMe-Backend/api/config"
	MenuImplRepositories "GymMe-Backend/api/repositories/menu/repositories-menu-impl"
	"GymMe-Backend/api/service/menu"
	menuserviceimpl "GymMe-Backend/api/service/menu/menu-service-impl"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func setupBookmark() (*gorm.DB, menu.BookmarkService) {
	configenv.InitEnvConfigs(true, "")
	db := configenv.InitDB()
	BookmarkRepo := MenuImplRepositories.NewBookmarkRepositoryImpl()
	Bookmarkservice := menuserviceimpl.NewBookmarkServiceImpl(db, BookmarkRepo)
	return db, Bookmarkservice
}

func TestInsertBookmark(t *testing.T) {
	db, service := setupBookmark()
	userId := 1
	InformationId := 3
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	_, err := service.AddBookmark(userId, InformationId)
	if err != nil {
		t.Errorf("Failed On: %v", err)
	}
	t.Log("Bookmark Inserted Successfully")
	assert.Nil(t, err, nil)

}
