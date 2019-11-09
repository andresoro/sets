package main

// StructSet implements a set using empty struct as value
type StructSet map[interface{}]struct{}

// BoolSet implements set using bool as value
type BoolSet map[interface{}]bool

// InterfaceSet implements a set using an interface as value
type InterfaceSet map[interface{}]interface{}

func main() {

}
