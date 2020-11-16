package main

import "cattleai/ent"

func main() {
	GenCates("PregnancyTestType")
	GenCates("PregnancyTestMethod")
	GenCates("PregnancyTestResult")
	Gen("PregnancyTest", &ent.PregnancyTestCreate{})
}
