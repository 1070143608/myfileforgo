package house

import "github.com/bigwhite/functrace"

type House struct {
	Material     string
	HasFireplace bool
	Floors       int
}

func NewHouse(opts ...HouseOption) *House {
	defer functrace.Trace()()
	const (
		defaultFloors       = 2
		defaultHasFireplace = true
		defaultMaterial     = "wood"
	)

	h := &House{
		Material:     defaultMaterial,
		HasFireplace: defaultHasFireplace,
		Floors:       defaultFloors,
	}
	for _, opt := range opts {
		// Call the option giving the instantiated
		// *House as the argument
		opt(h)
	}
	return h
}

type HouseOption func(*House)

func WithConcrete() HouseOption {
	defer functrace.Trace()()
	return func(h *House) {
		h.Material = "concrete"
	}
}

func WithoutFireplace() HouseOption {
	defer functrace.Trace()()
	return func(h *House) {
		h.HasFireplace = false
	}
}

func WithFloors(floors int) HouseOption {
	defer functrace.Trace()()
	return func(h *House) {
		h.Floors = floors
	}
}
