package db

import (
	"github.com/0glabs/0g-serving-agent/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DB struct {
	db *gorm.DB
}

func NewDB(conf *config.Config) (*DB, error) {
	db, err := gorm.Open(mysql.Open(conf.Database.Provider), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}
