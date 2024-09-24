package migrate

import (
	configenv "GymMe-Backend/api/config"
	"GymMe-Backend/api/entities"
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
	dsn := fmt.Sprintf("sqlserver://%s:%d?database=%s&connection+timeout=30&encrypt=disable&trustServerCertificate=false&app name=SqlClient",
		configenv.EnvConfigs.Hostname,
		configenv.EnvConfigs.DBPort,
		configenv.EnvConfigs.DBName)
	fmt.Println(dsn)

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
	var tableNames []string
	db.Raw("SELECT table_name FROM information_schema.tables WHERE table_type = 'BASE TABLE'").Scan(&tableNames)
	//
	//// Drop each table
	for _, tableName := range tableNames {
		err := db.Migrator().DropTable(tableName)
		if err != nil {
			log.Printf("Failed to drop table %s: %v", tableName, err)
		} else {
			log.Printf("Successfully dropped table %s", tableName)
		}
	}
	//migrate all table if neccesay
	//err = database.DropAllDatabase(db)
	//if err != nil {
	//	log.Printf("%s Failed to drop all tables with error: %s", logEntry, err)
	//	panic(err)
	//}
	//log.Printf("%s Successfully dropped all tables", logEntry)

	//end off drop all table
	err = db.AutoMigrate(
		&entities.Users{},
		&entities.UserDetail{},
		&entities.Users{},
		&entities.PaymentMethod{},
		&entities.InformationType{},
		&entities.Information{},
		&entities.BookmarkType{},
		&entities.Bookmark{},
	)

	if err != nil {
		log.Printf("%s Failed with error: %s", logEntry, err)
		panic(err)
	}

	log.Printf("%s Success", logEntry)
}
