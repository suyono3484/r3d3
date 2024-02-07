package r3d3

//go:generate mockgen -package r3d3 -destination types_mock.go . Armament,SpaceCraft

type Status interface {
	String() string
}

type Armament interface {
	Title() string
	Qty() int
}

type SpaceCraftInList interface {
	ID() int64
	Name() string
	Status() string
}

type SpaceCraft interface {
	ID() int64
	Name() string
	Class() string
	Crew() uint64
	ImageURL() string
	Value() float64
	Status() string
	Armament() []Armament
}

type ListFilter interface {
	Param() string
}

type ArmamentSpaceCraftCreate interface {
	ID() int64
	Qty() int
}

type SpaceCraftCreate interface {
	Name() string
	Class() string
	Crew() uint64
	ImageURL() string
	Value() float64
	Status() string
	Armament() []ArmamentSpaceCraftCreate
}
