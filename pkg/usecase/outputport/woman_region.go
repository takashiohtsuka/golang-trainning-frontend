package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/domain/collection"
	"golang-trainning-frontend/pkg/domain/entity"
)

type WomanRegionRepository interface {
	FindPickupByRegion(ctx context.Context, regionID uint) (collection.Collection[entity.WomanEntity], error)
}
