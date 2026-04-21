package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/querymodel"
)

type WomanRegionRepository interface {
	FindPickupByRegion(ctx context.Context, regionID uint) (collection.Collection[querymodel.WomanQueryModel], error)
}
