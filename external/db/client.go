package db

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type DbClient struct {
	*gorm.DB
}

func NewDbClient() *DbClient {
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_URL")), &gorm.Config{})
	if err != nil {
		logrus.Fatal("Could not connect to postgres")
	}

	return &DbClient{db}
}
