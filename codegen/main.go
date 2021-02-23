package main

import "cattleai/ent"

func main() {
	// GenCates("EpidemicType")
	Gen("ShedSetting", &ent.ShedSettingCreate{})
	Gen("WarehouseSetting", &ent.WarehouseSettingCreate{})
	Gen("Customer", &ent.CustomerCreate{})
}
