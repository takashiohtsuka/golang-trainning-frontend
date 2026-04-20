package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/dto"
)

type StoreRegionRepository interface {
	FindPickupByRegion(ctx context.Context, regionID uint) (collection.Collection[dto.StoreDTO], error)
}
