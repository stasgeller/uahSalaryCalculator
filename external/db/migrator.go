package db

import (
	"uahSalaryBot/infrastructure/domain"

	"github.com/sirupsen/logrus"
)

func main() {
	client := NewDbClient()
	if err := client.AutoMigrate(&domain.User{}); err != nil {
		logrus.Errorf("[migration]: %s", err.Error())
	}
}
