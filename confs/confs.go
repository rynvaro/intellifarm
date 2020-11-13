package confs

var AppConfs = &Confs{
	InitConf: &InitConf{},
}

type Confs struct {
	InitConf *InitConf `json:"initConf"`
}

type InitConf struct {
	FarmCategoryInitialized  bool `json:"farmCategoryInitialized"`
	ShedCategoryInitialized  bool `json:"shedCategoryInitialized"`
	ShedTypeInitialized      bool `json:"shedTypeInitialized"`
	BreathRateInitialized    bool `json:"breathRateInitialized"`
	HairStateInitialized     bool `json:"hairStateInitialized"`
	WindDirectionInitialized bool `json:"windDirectionInitialized"`
	PositionInitialized      bool `json:"positionInitialized"`
	DutyInitialized          bool `json:"dutyInitialized"`
}
