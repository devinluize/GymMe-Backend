package migration

import (
	configenv "GymMe-Backend/api/config"
	entities "GymMe-Backend/api/entities/auth"
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

	dsn := "sqlserver://localhost:1433?database=assignmentDB&connection+timeout=30&encrypt=disable&trustServerCertificate=false&app name=SqlClient"
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
	err = db.AutoMigrate(
		&entities.RegisterPayloads{},
	)

	if err != nil {
		log.Printf("%s Failed with error: %s", logEntry, err)
		panic(err)
	}

	log.Printf("%s Success", logEntry)
}
