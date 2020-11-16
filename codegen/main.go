package main

import "cattleai/ent"

func main() {
	GenCates("EstrusType")
	Gen("Estrus", &ent.EstrusCreate{})
}
