package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/querymodel"
)

type StorePrefectureRepository interface {
	FindAllByPrefecture(ctx context.Context, prefectureID uint) (collection.Collection[querymodel.StoreQueryModel], error)
}
