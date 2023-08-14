package interfaces

import "github.com/jeanSagaz/server/internal/domain/entity"

type IUsdBrlRepository interface {
	Insert(customer *entity.UsdBrl) (*entity.UsdBrl, error)
}
