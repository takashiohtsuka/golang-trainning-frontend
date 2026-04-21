package women

import "golang-trainning-frontend/pkg/usecase/input"

type DetailRequest struct {
	ID uint `param:"id" validate:"required"`
}

func (req *DetailRequest) ToInput() input.GetWomanDetailInput {
	return input.GetWomanDetailInput{
		WomanID: req.ID,
	}
}
