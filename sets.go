package sets

import (
	mapset "github.com/deckarep/golang-set"
)

// StructSet implements a set using empty struct as value
type StructSet map[interface{}]struct{}

// BoolSet implements set using bool as value
type BoolSet map[interface{}]bool

// InterfaceSet implements a set using an interface as value
type InterfaceSet map[interface{}]interface{}

// MapSet is imported set from deckarep/golang-set
type MapSet mapset.Set
