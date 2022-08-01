package helper

import (
	"github.com/sirupsen/logrus"
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabaseTesting() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		defer logrus.Info("Database connection failed")
		logrus.Fatal(err)
		return nil
	}

	//  Initialize all model for auto migration here
	err = db.AutoMigrate(
		&models.Category{},
		&models.Product{},
		&models.Outlet{},
		&models.OutletProduct{},
		&models.User{},
	)

	if err != nil {
		defer logrus.Info("Database migration failed")
		logrus.Fatal(err)
		return nil
	}

	return db
}
