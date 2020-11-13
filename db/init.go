package db

import (
	"cattleai/confs"
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

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
}

func dataInit() {
	conf := confs.AppConfs
	ctx := context.Background()
	// get init config
	confs, err := Client.Conf.Query().All(ctx)
	if err != nil {
		panic(err)
	}
	if len(confs) > 0 {
		conf = confs[0].Confs
	}
	// init farm category
	if !conf.InitConf.FarmCategoryInitialized {
		for _, v := range initData["categories"] {
			_, err := Client.Category.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("category %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("category %s initialized", v))
		}
		conf.InitConf.FarmCategoryInitialized = true
	}

	// init shed category
	if !conf.InitConf.ShedCategoryInitialized {
		for _, v := range initData["shedCategories"] {
			_, err := Client.ShedCategory.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("shedCategory %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("shedCategory %s initialized", v))
		}
		conf.InitConf.ShedCategoryInitialized = true
	}

	// init shed type
	if !conf.InitConf.ShedTypeInitialized {
		for _, v := range initData["shedTypes"] {
			_, err := Client.ShedType.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("shedType %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("shedType %s initialized", v))
		}
		conf.InitConf.ShedTypeInitialized = true
	}

	// init breath rates
	if !conf.InitConf.BreathRateInitialized {
		for _, v := range initData["breathRates"] {
			_, err := Client.BreathRate.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("breathRate %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("breathRate %s initialized", v))
		}
		conf.InitConf.BreathRateInitialized = true
	}

	// init hair states
	if !conf.InitConf.HairStateInitialized {
		for _, v := range initData["hairStates"] {
			_, err := Client.HairState.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("hairState %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("hairState %s initialized", v))
		}
		conf.InitConf.HairStateInitialized = true
	}

	// init wind directions
	if !conf.InitConf.WindDirectionInitialized {
		for _, v := range initData["windDirections"] {
			_, err := Client.WindDirection.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("windDirection %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("windDirection %s initialized", v))
		}
		conf.InitConf.WindDirectionInitialized = true
	}

	// init positions
	if !conf.InitConf.PositionInitialized {
		for _, v := range initData["positions"] {
			_, err := Client.Position.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("position %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("position %s initialized", v))
		}
		conf.InitConf.PositionInitialized = true
	}

	// init positions
	if !conf.InitConf.DutyInitialized {
		for _, v := range initData["duties"] {
			_, err := Client.Duty.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("duty %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("duty %s initialized", v))
		}
		conf.InitConf.DutyInitialized = true
	}

	// update init config
	if len(confs) > 0 {
		_, err := Client.Conf.UpdateOneID(confs[0].ID).SetConfs(conf).Save(ctx)
		if err != nil {
			log.Error().Msg(err.Error())
		}
	} else { // or add new
		_, err := Client.Conf.Create().SetConfs(conf).Save(ctx)
		if err != nil {
			log.Error().Msg(err.Error())
		}
	}
}
