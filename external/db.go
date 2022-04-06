package external

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbClient struct {
	*gorm.DB
}

func NewDbClient() *DbClient {
	db, err := gorm.Open(sqlite.Open("./db/salary_bot"), &gorm.Config{})
	if err != nil {
		logrus.Fatal("Could not connect to sqlite")
	}

	return &DbClient{db}
}
