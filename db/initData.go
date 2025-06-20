package db

var initData = map[string][]string{
	"categories": {
		"繁育场",
		"育肥场",
	},
	"shedCategories": {
		"封闭式",
		"半封闭式",
		"开放式",
	},
	"shedTypes": {
		"犊牛舍",
		"后备母牛舍",
		"繁育母牛舍",
		"育肥牛舍",
		"隔离牛舍",
	},
	"breathRates": {
		"极速浅呼吸",
		"开口呼吸",
		"无明显表现",
	},
	"hairStates": {
		"干燥清洁",
		"部分污物",
		"腿上污物",
		"大量污物",
	},
	"windDirections": {
		"东",
		"东南",
		"南",
		"西南",
		"西",
		"西北",
		"北",
		"东北",
	},
	"positions": {
		"场长/经理",
		"副厂长/副经理",
		"技术负责人",
		"组长/班长",
		"员工/工人",
	},
	"duties": {
		"营养师",
		"繁育师",
		"兽医师",
		"信息员",
		"饲养员",
		"库管员",
		"其他",
	},
	"cattleGenders": {
		"母牛",
		"公牛",
		"阉牛",
	},
	"cattleCates": {
		"育肥牛/1,2,3",
		"繁育母牛/1",
		"种用公牛/2",
	},
	"reproductiveStates": {
		"出生",
		"配种",
		"怀孕",
		"空怀",
	},
	"cattleTypes": {
		"杂交",
		"纯种",
	},
	"cattleJoinedTypes": {
		"自繁",
		"购入",
		"转入",
	},
	"cattleHairColors": {
		"黑白",
		"红白",
		"全黑",
		"全红",
	},
	"breedingTypes": {
		"人工授精",
		"胚胎移植",
		"自然受精",
	},
	"cattleOwners": {
		"本场牛只",
		"产业扶贫",
		"外部寄养",
		"其他",
	},
	"calveTypes": {
		"自然分娩",
		"人工助产",
		"器械助产",
		"难产",
		"破腹产",
	},
	"calveCounts": {
		"单母胎",
		"单公胎",
		"双母胎",
		"双公胎",
		"异性双胎",
		"其他",
	},
	"estrusTypes": {
		"自然发情",
		"同期发情",
	},
	"semenFrozenTypes": {
		"常规",
		"后测常规",
		"性控",
		"后测性控",
		"胚胎",
	},
	"pregnancyTestTypes": {
		"初检",
		"复检",
	},
	"pregnancyTestMethods": {
		"直肠检查",
		"B超检查",
		"试剂盒检查",
		"试纸条检查",
	},
	"pregnancyTestResults": {
		"怀孕",
		"空怀",
		"空怀待复检",
	},
	"abortionTypes": {
		"干胎",
		"鲜胎",
		"其他",
	},
	"treatmentResults": {
		"已痊愈",
		"死亡",
		"淘汰",
	},
	"treatmentStates": {
		"治疗中",
		"治疗结束",
	},
	"whereablouts": {
		"转入牛群",
		"淘汰屠宰",
		"无害处理",
	},
}
