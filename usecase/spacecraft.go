package usecase

import "github.com/suyono3484/r3d3"

type Model interface {
	GetByID(int64) r3d3.SpaceCraft
}

type UseCase struct {
	model Model
}

func NewUseCase(model Model) *UseCase {
	return &UseCase{
		model: model,
	}
}

func (u *UseCase) List( /* need some filter */ ) {

}

func (u *UseCase) Get(id uint64) {

}
