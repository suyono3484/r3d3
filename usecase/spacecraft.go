package usecase

import "github.com/suyono3484/r3d3"

type Model interface {
	GetByID(int64) (r3d3.SpaceCraft, error)
	List(filter ...r3d3.ListFilter) ([]r3d3.SpaceCraftInList, error)
	Create(r3d3.SpaceCraftCreate) error
	Update(int64, r3d3.SpaceCraftCreate, []string) error
	Delete(int64) error
}

type UseCase struct {
	model Model
}

func NewUseCase(model Model) *UseCase {
	return &UseCase{
		model: model,
	}
}

func (u *UseCase) List(filter ...r3d3.ListFilter) ([]r3d3.SpaceCraftInList, error) {
	//TODO: this is a placeholder, update me!
	return u.model.List(filter...)
}

func (u *UseCase) Get(id int64) (r3d3.SpaceCraft, error) {
	//TODO: this is a placeholder, update me!
	return u.model.GetByID(id)
}

func (u *UseCase) Create(spacecraft r3d3.SpaceCraftCreate, fields []string) error {
	//TODO: this is a placeholder, update me!
	return u.model.Create(spacecraft)
}

func (u *UseCase) Update(id int64, spacecraft r3d3.SpaceCraftCreate, fields []string) error {
	//TODO: this is a placeholder, update me!
	return u.model.Update(id, spacecraft, fields)
}

func (u *UseCase) Delete(id int64) error {
	//TODO: this is a placeholder, update me!
	return u.model.Delete(id)
}
