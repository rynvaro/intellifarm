package main

import "cattleai/ent"

func main() {
	GenCates("SemenFrozenType")
	Gen("Breeding", &ent.BreedingCreate{})
}
