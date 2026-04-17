package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/domain/collection"
	"golang-trainning-frontend/pkg/domain/entity"
)

type StorePrefectureRepository interface {
	FindAllByPrefecture(ctx context.Context, prefectureID uint) (collection.Collection[entity.StoreEntity], error)
}
