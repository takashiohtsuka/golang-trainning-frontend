package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/querymodel"
)

type WomanPrefectureRepository interface {
	FindAllByPrefecture(ctx context.Context, prefectureID uint) (collection.Collection[querymodel.WomanQueryModel], error)
}
