package config

import (
	"fmt"

	"github.com/DeniesKresna/brigunaduty/types/constants"
	"github.com/DeniesKresna/gohelper/utstring"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (cfg *Config) SetConfigDatabase() (err error) {
	user := utstring.GetEnv(constants.ENV_DB_USER)
	pass := utstring.GetEnv(constants.ENV_DB_PASSWORD)
	host := utstring.GetEnv(constants.ENV_DB_HOST)
	port := utstring.GetEnv(constants.ENV_DB_PORT)
	dbname := utstring.GetEnv(constants.ENV_DB_NAME)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	if utstring.GetEnv(constants.ENV_APP_ENV, "local") != "production" {
		cfg.DB = db.Debug()
	} else {
		cfg.DB = db
	}

	return
}
