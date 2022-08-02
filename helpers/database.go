package helpers

import (
	"github.com/sirupsen/logrus"
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var MYSQL_URL = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
var MYSQL_USER = "root"
var MYSQL_PASSWORD = "bismillah"
var MYSQL_HOST = "localhost"
var MYSQL_PORT = "3306"
var MYSQL_DB = "mini_pos_go"

func SetupDatabaseTesting() *gorm.DB {
	// db, err := gorm.Open(
	// 	mysql.Open(
	// 		fmt.Sprintf(
	// 			MYSQL_URL,
	// 			MYSQL_USER,
	// 			MYSQL_PASSWORD,
	// 			MYSQL_HOST,
	// 			MYSQL_PORT,
	// 			MYSQL_DB,
	// 		)), &gorm.Config{})

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

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
