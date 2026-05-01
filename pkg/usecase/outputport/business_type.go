package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/querymodel"
)

type BusinessTypeRepository interface {
	FindAll(ctx context.Context) ([]querymodel.BusinessTypeQueryModel, error)
}
