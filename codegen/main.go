package main

import "cattleai/ent"

func main() {
	// GenCates("EpidemicType")
	Gen("Event", &ent.EventCreate{})
	Gen("Operation", &ent.OperationCreate{})
}
