package main

import "cattleai/ent"

func main() {
	GenCates("AbortionType")
	Gen("Abortion", &ent.AbortionCreate{})
}
