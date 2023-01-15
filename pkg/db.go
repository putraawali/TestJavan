package pkg

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testjavan/helpers/constants"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	showRecordNotFound := false

	if os.Getenv("ENV") == "DEVELOPMENT" {
		showRecordNotFound = true
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,        // Slow SQL threshold
			LogLevel:                  logger.Error,       // Log level
			IgnoreRecordNotFoundError: showRecordNotFound, // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,              // Disable color
		},
	)

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	pgDB, err := connection.DB()
	if err != nil {
		return nil, err
	}

	if _, found := os.LookupEnv("DB_MAX_OPEN_CONNECTION"); found {
		if maxOpenConnection, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTION")); err == nil {
			pgDB.SetMaxOpenConns(maxOpenConnection)
		}
	}

	if _, found := os.LookupEnv("DB_MAX_IDLE_CONNECTION"); found {
		if maxIdleConnection, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTION")); err == nil {
			pgDB.SetMaxIdleConns(maxIdleConnection)
		}
	}

	if _, found := os.LookupEnv("DB_MAX_LIFETIME"); found {
		if maxLifeTime, err := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME")); err == nil {
			pgDB.SetConnMaxLifetime(time.Duration(maxLifeTime) * time.Minute)
		}
	}

	fmt.Printf("[%s] Success connect to database\n", time.Now().Format(constants.TimeFormat))
	return connection, nil
}
