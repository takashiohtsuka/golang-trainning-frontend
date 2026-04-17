package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/domain/collection"
	"golang-trainning-frontend/pkg/domain/entity"
)

type WomanDistrictRepository interface {
	FindAllByDistrict(ctx context.Context, districtID uint) (collection.Collection[entity.WomanEntity], error)
}
