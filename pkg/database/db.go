package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ribice/gorsk/pkg/utl/config"
)

// CreateConnection - ...
func CreateConnection() (*gorm.DB, error) {
	host := config.Database.Host
	databaseUser := config.Database.User
	databaseName := config.Database.Name
	databasePassword := config.Database.Password

	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=disable password=%s",
			host, databaseUser, databaseName, databasePassword,
		),
	)
}
