package r3d3

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
