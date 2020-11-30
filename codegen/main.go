package main

import "cattleai/ent"

func main() {
	// GenCates("EpidemicType")
	Gen("CattleIn", &ent.CattleInCreate{})
	Gen("CattleOut", &ent.CattleOutCreate{})
}
