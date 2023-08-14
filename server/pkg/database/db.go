package database

import (
	"fmt"

	"github.com/jeanSagaz/server/internal/domain/entity"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Db            *gorm.DB
	AutoMigrateDb bool
}

func NewDb() *Database {
	return &Database{}
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	// d.Db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		"sa",
		"SqlServer2019!",
		"localhost",
		1434,
		"fc")
	d.Db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	if d.AutoMigrateDb {
		err = d.Db.Table("UsdBrl").AutoMigrate(&entity.UsdBrl{})
		if err != nil {
			panic(err)
		}
	}

	return d.Db, nil
}

func GetTableName(tableName string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table(tableName)
	}
}
