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

// Add element to StructSet
func (s StructSet) Add(e interface{}) {
	s[e] = struct{}{}
}

// Has element in StructSet
func (s StructSet) Has(e interface{}) bool {
	if _, ok := s[e]; ok {
		return true
	}
	return false
}

// Add element to BoolSet
func (s BoolSet) Add(e interface{}) {
	s[e] = true
}

// Has element in BoolSet
func (s BoolSet) Has(e interface{}) bool {
	if _, ok := s[e]; ok {
		return true
	}
	return false
}

// Add element to InterfaceSet
func (s InterfaceSet) Add(e interface{}) {
	s[e] = true
}

// Has element in InterfaceSet
func (s InterfaceSet) Has(e interface{}) bool {
	if _, ok := s[e]; ok {
		return true
	}
	return false
}
