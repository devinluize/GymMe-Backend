package migrate

import (
	configenv "GymMe-Backend/api/config"
	entities "GymMe-Backend/api/entities/Equipment"
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

func Migrate() {
	configenv.InitEnvConfigs(false, "")
	logEntry := "Auto Migrating to database"
	//dsn := "Server=localhost\\MSSQLSERVER01;Database=GymMe;Trusted_Connection=True;"
	//dsn := fmt.Sprintf("sqlserver://34.101.163.215:%d?database=%s&connection+timeout=30&encrypt=disable&trustServerCertificate=false&app name=SqlClient",
	//	configenv.EnvConfigs.DBPort,
	//	configenv.EnvConfigs.DBName)
	//fmt.Println(dsn)
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&connection+timeout=30&encrypt=disable&trustServerCertificate=false&app name=SqlClient",
		configenv.EnvConfigs.DBUser,
		configenv.EnvConfigs.DBPass,
		configenv.EnvConfigs.DBHost,
		configenv.EnvConfigs.DBPort,
		configenv.EnvConfigs.DBName)

	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	if err != nil {
		log.Printf("%s Failed to connect to database with error: %s", logEntry, err)
		panic(err)
	}
	// Get the list of all tables
	fmt.Println("check point")
	var tableNames []string
	db.Raw("SELECT table_name FROM information_schema.tables WHERE table_type = 'BASE TABLE'").Scan(&tableNames)
	//
	//// Drop each table
	//for _, tableName := range tableNames {
	//	err := db.Migrator().DropTable(tableName)
	//	if err != nil {
	//		log.Printf("Failed to drop table %s: %v", tableName, err)
	//	} else {
	//		log.Printf("Successfully dropped table %s", tableName)
	//	}
	//	//}
	//}
	/////
	//migrate all table if neccesay
	//err = database.DropAllDatabase(db)
	//if err != nil {
	//	log.Printf("%s Failed to drop all tables with error: %s", logEntry, err)
	//	panic(err)
	//}
	//log.Printf("%s Successfully dropped all tables", logEntry)

	//end off drop all table
	err = db.AutoMigrate(
		//&entities.Users{},
		//&entities.UserDetail{},
		//&entities.Users{},
		//&entities.PaymentMethod{},
		//&entities.ArticleType{},
		//&entities.ArticleEntities{},
		//&entities.BookmarkType{},
		//&entities.Bookmark{},
		//&entities.ArticleBodyEntities{},
		//&entities.WeightHistoryEntities{},
		//&entities.CalendarEntity{},
		//&entities.WeightHistoryEntities{},
		//&entities.TimerEntity{},
		////-----
		//&entities.TimerQueueEntity{},
		////equipment
		//&entities2.EquipmentDetailEntity{},
		//&entities2.EquipmentDetailEntity{},
		&entities.EquipmentDifficultyEntities{},
		//
		////$entities2.EquipmentMasterEntities{}
		//&entities2.EquipmentMasterEntities{},
		&entities.EquipmentTypeEntity{},
		&entities.ForceTypeEntities{},
		&entities.MuscleGroupEntities{},
		//&entities2.EquipmentProfileEntity{},
		//
		//&entities.EquipmentDetailEntity{},

		//------
		&entities.EquipmentCourseDataEntity{},
		//
		//&entities.EquipmentCourseDataEntity{},
		//&entities2.EquipmentBookmark{},
		//&entities2.SearchHistoryEntities{},
		//&entities.EquipmentSearchHistoryEntities{},
	)

	//&entities.EquipmentTypeEntity{},
	//&entities.ForceTypeEntities{},
	//&entities.MuscleGroupEntities{},
	//&entities.EquipmentDifficultyEntities{},
	//&entities.EquipmentMasterEntities{},
	//&entities.EquipmentDetailEntity{},
	if err != nil {
		log.Printf("%s Failed with error: %s", logEntry, err)
		panic(err)
	}

	log.Printf("%s Success", logEntry)
}
