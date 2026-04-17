package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/domain/collection"
	"golang-trainning-frontend/pkg/domain/entity"
)

type StoreRegionRepository interface {
	FindPickupByRegion(ctx context.Context, regionID uint) (collection.Collection[entity.StoreEntity], error)
}
