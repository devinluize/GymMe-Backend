package MenuImplRepositories

import (
	"GymMe-Backend/api/entities"
	"GymMe-Backend/api/helper"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	menuRepository "GymMe-Backend/api/repositories/menu"
	"errors"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type InformationMenu struct {
}

func NewInformationMenu() menuRepository.InformationMenu {
	return &InformationMenu{}
}
func (i *InformationMenu) UpdateInformation(tx *gorm.DB, payloads MenuPayloads.InformationUpdatePayloads) (entities.InformationEntities, *responses.ErrorResponses) {
	var InformationEntities entities.InformationEntities
	var InformationBodyEntities []entities.InformationBodyEntities

	err := tx.Model(&InformationEntities).Where(entities.InformationEntities{InformationId: payloads.InformationId}).First(&InformationEntities).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return InformationEntities, &responses.ErrorResponses{
				Message:    err.Error(),
				StatusCode: http.StatusNotFound,
				Err:        err,
			}
		}
		return InformationEntities,
			&responses.ErrorResponses{StatusCode: http.StatusInternalServerError,
				Err:     err,
				Message: err.Error()}
	}
	for _, i := range payloads.InformationBodyContent {
		InformationBodyEntitiesData := entities.InformationBodyEntities{
			InformationBodyParagraph:    i.InformationBodyParagraph,
			InformationImageContentPath: i.InformationImageContentPath,
			InformationId:               payloads.InformationId,
		}
		InformationBodyEntities = append(InformationBodyEntities, InformationBodyEntitiesData)
	}
	//Information Body Inserting
	InformationEntities.InformationBody = InformationBodyEntities
	err = tx.Delete(entities.InformationBodyEntities{}, entities.InformationBodyEntities{InformationId: payloads.InformationId}).Error
	err = tx.Save(&InformationEntities).Error
	if err != nil {
		return InformationEntities, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			//Err:        nil,
			Success: false,
		}
	}
	return InformationEntities, nil
}
func (i *InformationMenu) DeleteInformationById(db *gorm.DB, id int) (bool, *responses.ErrorResponses) {
	err := db.Model(&entities.InformationEntities{}).Where(entities.InformationEntities{InformationId: id}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, &responses.ErrorResponses{
				StatusCode: http.StatusBadRequest,
				Message:    "Delete Failed Id Not Found",
				Success:    false,
			}
		}
		return false, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	err = db.Delete(entities.InformationBodyEntities{}, entities.InformationBodyEntities{InformationId: id}).Error
	if err != nil {
		return false, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}

	err = db.Delete(entities.InformationEntities{}, entities.InformationEntities{InformationId: id}).Error
	if err != nil {
		return false, &responses.ErrorResponses{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}
	return true, nil
}
func (i *InformationMenu) InsertInformation(tx *gorm.DB, payloads MenuPayloads.InformationInsertPayloads) (entities.InformationEntities, *responses.ErrorResponses) {
	var Entities entities.InformationEntities
	Entities = entities.InformationEntities{
		//InformationId:                0,
		InformationHeader:            payloads.InformationHeader,
		InformationHeaderPathContent: payloads.InformationHeaderPathContent,
		//InformationImageContentPath1: payloads.InformationImageContentPath1,
		//InformationImageContentPath2: payloads.InformationImageContentPath2,
		//InformationImageContentPath3: payloads.InformationImageContentPath3,
		//InformationImageContentPath4: payloads.InformationImageContentPath4,
		//InformationImageContentPath5: payloads.InformationImageContentPath5,
		//InformationTypeId:      payloads.InformationTypeId,
		InformationDateCreated: time.Now(),
	}
	//for _, detail := range payloads.InformationBodyParagraph {
	//	Entities.InformationBody = append(Entities.InformationBody, entities.InformationBodyEntities{
	//		InformationBodyParagraph:    detail.InformationBodyParagraph,
	//		InformationImageContentPath: detail.InformationImageContentPath,
	//	})
	//}
	var EntitiesDetail []entities.InformationBodyEntities

	//Entities.InformationId = Entities.InformationId
	err := tx.Create(&Entities).First(&Entities).Error
	if err != nil {
		return Entities, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Err:        err,
			Success:    false,
		}
	}

	//err = tx.Model(&Entities).Where(entities.InformationEntities{InformationId: Entities.InformationId}).First(&Entities).Error
	for _, paragraph := range payloads.InformationBodyParagraph {
		EntitiesDetail = append(EntitiesDetail, entities.InformationBodyEntities{
			InformationImageContentPath: paragraph.InformationImageContentPath,
			InformationBodyParagraph:    paragraph.InformationBodyParagraph,
			InformationId:               Entities.InformationId,
		})
	}
	err = tx.Create(&EntitiesDetail).Error
	if err != nil {
		return Entities, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Err:        err,
			Success:    false,
		}
	}
	return Entities, nil
}
func (i *InformationMenu) GetInformationById(db *gorm.DB, id int, userId int, cld *cloudinary.Cloudinary) (MenuPayloads.InformationSelectResponses, *responses.ErrorResponses) {
	EntitiesInfo := entities.InformationEntities{}
	err := db.Model(&entities.InformationEntities{}).Where(entities.InformationEntities{InformationId: id}).First(&EntitiesInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return MenuPayloads.InformationSelectResponses{}, &responses.ErrorResponses{
				StatusCode: http.StatusBadRequest,
				Message:    "Delete Failed Id Not Found",
			}
		}
	}
	ResDetail := []MenuPayloads.InformationBodyDetail{}
	err = db.Model(&entities.InformationBodyEntities{}).Where(entities.InformationBodyEntities{InformationId: id}).Scan(&ResDetail).Error
	if err != nil {
		return MenuPayloads.InformationSelectResponses{}, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}
	SelectPayload := []MenuPayloads.InformationBodyDetail{}
	for _, detail := range ResDetail {
		if detail.InformationImageContentPath != "" {
			urls, _ := cld.Image(detail.InformationImageContentPath)
			//res.SortOf = url
			detail.InformationImageContentPath = fmt.Sprintf("https://res.cloudinary.com/%s/%s/%s/%s",
				"dlrd9z1mk",          // Replace with your Cloudinary cloud name
				urls.AssetType,       // e.g., "image"
				urls.DeliveryType,    // e.g., "upload"
				urls.PublicID+".jpg", // Add appropriate file extension
			)
		}
		SelectPayloadData := MenuPayloads.InformationBodyDetail{
			InformationBodyParagraph:    detail.InformationBodyParagraph,
			InformationImageContentPath: detail.InformationImageContentPath,
		}
		SelectPayload = append(SelectPayload, SelectPayloadData)
	}
	isBookmarkExist := false
	urls, _ := cld.Image(EntitiesInfo.InformationHeaderPathContent)
	//res.SortOf = url
	EntitiesInfo.InformationHeaderPathContent = fmt.Sprintf("https://res.cloudinary.com/%s/%s/%s/%s",
		"dlrd9z1mk",          // Replace with your Cloudinary cloud name
		urls.AssetType,       // e.g., "image"
		urls.DeliveryType,    // e.g., "upload"
		urls.PublicID+".jpg", // Add appropriate file extension
	)

	//EntitiesInfo.InformationHeaderPathContent
	err = db.Model(&entities.Bookmark{}).Where(entities.Bookmark{InformationId: id, UserId: userId}).Select("1").
		Scan(&isBookmarkExist).Error
	result := MenuPayloads.InformationSelectResponses{
		InformationHeader:            EntitiesInfo.InformationHeader,
		InformationDateCreated:       EntitiesInfo.InformationDateCreated,
		InformationCreatedByUserId:   EntitiesInfo.InformationCreatedByUserId,
		InformationId:                id,
		IsBookmark:                   isBookmarkExist,
		InformationBodyContent:       SelectPayload,
		InformationHeaderPathContent: EntitiesInfo.InformationHeaderPathContent,
		//InformationTypeId:          EntitiesInfo.InformationTypeId,
	}
	return result, nil
}
func (i *InformationMenu) GetAllInformationWithPagination(db *gorm.DB, paginationResponses helper.Pagination) (helper.Pagination, *responses.ErrorResponses) {

	var Entities []entities.InformationEntities
	//me := db.Model(&entities.InformationEntities{}) -> table joinan
	//cara 1
	//myJoinTable := db.Model(&entities.InformationEntities{}).
	//err := db.Model(&entities.InformationEntities{}).Scopes(database.Paginate(&Entities, &paginationResponses, me)).Order("information_id").Where("information_id <> 0").Scan(&Entities).Error
	//cara 2 langsung assign ke database nanti pilih aja apakah perlu buat join table atau ga kalau misalkan selectan itu merupakan hasil join table pake yang atas

	err := db.Model(&entities.InformationEntities{}).Scopes(helper.Paginate(&Entities, &paginationResponses, db)).Order("information_id").Where("information_id <> 0").Scan(&Entities).Error
	if err != nil {
		return paginationResponses, &responses.ErrorResponses{}
	}
	paginationResponses.Rows = Entities
	fmt.Println(paginationResponses.Rows)
	return paginationResponses, nil
}
func (i *InformationMenu) GetAllInformationWithFilter(db *gorm.DB, paginationResponses helper.Pagination, Key string, userId int, cloudinary *cloudinary.Cloudinary) (helper.Pagination, *responses.ErrorResponses) {
	//create history logging
	//ctx := context.Background()
	if Key != "" {

		historyLogging := entities.SearchHistoryEntities{
			UserId:     userId,
			SearchKey:  Key,
			DateSearch: time.Now(),
		}

		// Insert the new history record
		err := db.Create(&historyLogging).Error
		if err != nil {
			return paginationResponses, &responses.ErrorResponses{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
				Message:    "failed to log search history",
			}
		}

		// Check the total count of search history records for the user
		//historyCount := 0
		var historyCount int64
		errCount := db.Model(&entities.SearchHistoryEntities{}).
			Where("user_id = ?", userId).
			Count(&historyCount).Error
		if errCount != nil {
			return paginationResponses, &responses.ErrorResponses{
				StatusCode: http.StatusInternalServerError,
				Err:        errCount,
				Message:    "failed to check search history count",
			}
		}

		if historyCount > 10 {
			var excessCount int
			excessCount = int(historyCount) - 10
			// Delete oldest records if the count exceeds the limit
			var recordsToDelete []entities.SearchHistoryEntities
			errSelect := db.Where("user_id = ?", userId).
				Order("date_search ASC").
				Limit(excessCount).
				Find(&recordsToDelete).Error
			if errSelect != nil {
				return paginationResponses, &responses.ErrorResponses{
					StatusCode: http.StatusInternalServerError,
					Err:        errSelect,
					Message:    "failed to fetch old search history records",
				}
			}

			// Step 2: Delete the fetched records
			if len(recordsToDelete) > 0 {
				errDelete := db.Delete(&recordsToDelete).Error
				if errDelete != nil {
					return paginationResponses, &responses.ErrorResponses{
						StatusCode: http.StatusInternalServerError,
						Err:        errDelete,
						Message:    "failed to clean up old search history records",
					}
				}
			}
		}

		//	// Calculate how many records need to be deleted

		//	fmt.Println("excess : ")
		//	fmt.Println(excessCount)
		//	// Delete the oldest records
		//	errDelete := db.Where("user_id = ?", userId).
		//		Order("date_search ASC"). // Assuming `date_search` is used to track record age
		//		Limit(excessCount).
		//		Delete(&entities.SearchHistoryEntities{}).Error
		//	if errDelete != nil {
		//		return paginationResponses, &responses.ErrorResponses{
		//			StatusCode: http.StatusInternalServerError,
		//			Err:        errDelete,
		//			Message:    "failed to clean up old search history records",
		//		}
		//	}
	}

	var Entities []entities.InformationEntities
	//me := db.Model(&entities.InformationEntities{}) -> table joinan
	//cara 1
	joinTable := db.Model(&entities.InformationEntities{}).Where("information_id <> 0 AND information_header LIKE ? ", "%"+Key+"%")
	//err := db.Model(&entities.InformationEntities{}).Scopes(database.Paginate(&Entities, &paginationResponses, me)).Order("information_id").Where("information_id <> 0").Scan(&Entities).Error
	//cara 2 langsung assign ke database nanti pilih aja apakah perlu buat join table atau ga kalau misalkan selectan itu merupakan hasil join table pake yang atas
	err := joinTable.Scopes(helper.Paginate(&Entities, &paginationResponses, joinTable)).Order("information_id").Where("information_id <> 0 AND information_header LIKE ? ", "%"+Key+"%").Scan(&Entities).Error
	if err != nil {
		return paginationResponses, &responses.ErrorResponses{}
	}

	for i2, entity := range Entities {
		urls, _ := cloudinary.Image(entity.InformationHeaderPathContent)
		//res.SortOf = url
		Entities[i2].InformationHeaderPathContent = fmt.Sprintf("https://res.cloudinary.com/%s/%s/%s/%s",
			"dlrd9z1mk",          // Replace with your Cloudinary cloud name
			urls.AssetType,       // e.g., "image"
			urls.DeliveryType,    // e.g., "upload"
			urls.PublicID+".jpg", // Add appropriate file extension
		)
	}

	paginationResponses.Rows = Entities
	fmt.Println(paginationResponses.Rows)
	return paginationResponses, nil
}
func (i *InformationMenu) GetInformationHistory(db *gorm.DB, userId int) ([]entities.SearchHistoryEntities, *responses.ErrorResponses) {
	var entitiesData []entities.SearchHistoryEntities
	err := db.Model(&entities.SearchHistoryEntities{}).
		Where(entities.SearchHistoryEntities{UserId: userId}).
		Order("date_search DESC").
		Limit(10).
		Scan(&entitiesData).Error
	if err != nil {
		return entitiesData, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to get information history search data",
		}
	}
	return entitiesData, nil
}
