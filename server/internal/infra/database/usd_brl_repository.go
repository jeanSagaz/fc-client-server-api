package database

import (
	"context"
	"time"

	"github.com/jeanSagaz/server/internal/domain/entity"
	"github.com/jeanSagaz/server/pkg/database"
	"gorm.io/gorm"
)

type UsdBrlRepositoryDb struct {
	Db *gorm.DB
}

func NewUsdBrlRepositoryDb(db *gorm.DB) *UsdBrlRepositoryDb {
	return &UsdBrlRepositoryDb{Db: db}
}

func (repo UsdBrlRepositoryDb) Insert(usdBrl *entity.UsdBrl) (*entity.UsdBrl, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	err := repo.Db.WithContext(ctx).Scopes(database.GetTableName("UsdBrl")).Create(usdBrl).Error
	// err := repo.Db.Scopes(database.GetTableName("UsdBrl")).Create(usdBrl).Error
	if err != nil {
		return nil, err
	}

	return usdBrl, nil
}
