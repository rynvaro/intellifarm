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
	CattleGender             bool `json:"cattleGender"`
	CattleCate               bool `json:"cattleCate"`
	ReproductiveState        bool `json:"reproductiveState"`
	CattleType               bool `json:"cattleType"`
	CattleJoinedType         bool `json:"cattleJoinedType"`
	CattleHairColor          bool `json:"cattleHairColor"`
	BreedingType             bool `json:"breedingType"`
	CattleOwner              bool `json:"cattleOwner"`
	CalveType                bool `json:"calveType"`
	CalveCount               bool `json:"calveCount"`
	EstrusType               bool `json:"estrusType"`
	SemenFrozenType          bool `json:"semenFrozenType"`
	PregnancyTestType        bool `json:"pregnancyTestType"`
	PregnancyTestMethod      bool `json:"pregnancyTestMethod"`
	PregnancyTestResult      bool `json:"pregnancyTestResult"`
	AbortionType             bool `json:"abortionType"`
	TreatmentResult          bool `json:"treatmentResult"`
	TreatmentState           bool `json:"treatmentState"`
	Whereablout              bool `json:"whereabout"`
}
