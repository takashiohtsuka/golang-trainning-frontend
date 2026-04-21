package stores

import "golang-trainning-frontend/pkg/usecase/input"

type GetRequest struct {
	ID uint `param:"id" validate:"required"`
}

func (req *GetRequest) ToStoreDetailInput() input.GetStoreDetailInput {
	return input.GetStoreDetailInput{
		StoreID: req.ID,
	}
}

func (req *GetRequest) ToStoreWomanListInput() input.GetStoreWomanListInput {
	return input.GetStoreWomanListInput{
		StoreID: req.ID,
	}
}
