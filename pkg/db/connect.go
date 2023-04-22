package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(driver string, dsn string) (*gorm.DB, error) {
	var d gorm.Dialector
	switch driver {
	case "sqlite":
		d = sqlite.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported driver:" + driver)
	}

	return gorm.Open(d)
}
