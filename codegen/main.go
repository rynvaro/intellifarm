package main

import "cattleai/ent"

func main() {
	// GenCates("EpidemicType")
	// Gen("Event", &ent.EventCreate{})
	// Gen("Operation", &ent.OperationCreate{})
	// Gen("Medicine", &ent.MedicineCreate{})
	Gen("CattleIn", &ent.CattleInCreate{})
	Gen("CattleOut", &ent.CattleOutCreate{})
	Gen("CattleGroup", &ent.CattleGroupCreate{})
	Gen("CattleMove", &ent.CattleMoveCreate{})
	Gen("CattleDie", &ent.CattleDieCreate{})
}
