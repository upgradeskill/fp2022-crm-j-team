package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/pkg"
	"github.com/upgradeskill/fp2022-crm-j-team/routes"
)

func main() {

	/**
	* ========================
	*  Setup Application
	* ========================
	 */

	db := setupDatabase()
	app := setupApp()

	/**
	* ========================
	* Initialize All Route
	* ========================
	 */

	routes.Product(db, app)
	routes.NewRouteCategory(db, app)
	routes.NewRouteUser(db, app)

	/**
	* ========================
	*  Listening Server Port
	* ========================
	 */

	err := app.Start(":" + pkg.GodotEnv("PORT"))
	helpers.LogIfError(err, "Server is not running")

}

/**
* ========================
* Database Setup
* ========================
 */

func setupDatabase() *gorm.DB {
	db, err := gorm.Open(
		mysql.Open(
			fmt.Sprintf(
				pkg.GodotEnv("MYSQL_URL"),
				pkg.GodotEnv("MYSQL_USER"),
				pkg.GodotEnv("MYSQL_PASSWORD"),
				pkg.GodotEnv("MYSQL_HOST"),
				pkg.GodotEnv("MYSQL_PORT"),
				pkg.GodotEnv("MYSQL_DB"),
			)), &gorm.Config{})

	helpers.LogIfError(err, "Database connection failed")

	//  Initialize all model for auto migration here
	err = db.AutoMigrate(
		&models.Category{},
		&models.Product{},
		&models.Outlet{},
		&models.OutletProduct{},
		&models.User{},
	)

	helpers.LogIfError(err, "Database migration failed")

	// Import Seeder
	models.CreateUserSeeder(db)

	return db
}

/**
* ========================
* Application Setup
* ========================
 */

func setupApp() *echo.Echo {

	app := echo.New()

	return app
}
