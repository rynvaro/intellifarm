package db

import (
	"cattleai/confs"
	"context"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
)

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

	// init cattle genders
	if !conf.InitConf.CattleGender {
		for _, v := range initData["cattleGenders"] {
			_, err := Client.CattleGender.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("cattleGender %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("cattleGender %s initialized", v))
		}
		conf.InitConf.CattleGender = true
	}

	// init cattle cates
	if !conf.InitConf.CattleCate {
		for _, v := range initData["cattleCates"] {
			cate := strings.Split(v, "/")
			name := cate[0]
			genderIds := cate[1]
			_, err := Client.CattleCate.Create().SetName(name).SetGenderIds(genderIds).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("cattleCate %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("cattleCate %s initialized", v))
		}
		conf.InitConf.CattleCate = true
	}

	// init ReproductiveState
	if !conf.InitConf.ReproductiveState {
		for _, v := range initData["reproductiveStates"] {
			_, err := Client.ReproductiveState.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("reproductiveState %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("reproductiveState %s initialized", v))
		}
		conf.InitConf.ReproductiveState = true
	}

	// init CattleType
	if !conf.InitConf.CattleType {
		for _, v := range initData["cattleTypes"] {
			_, err := Client.CattleType.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("CattleType %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("CattleType %s initialized", v))
		}
		conf.InitConf.CattleType = true
	}

	// init CattleJoinedType
	if !conf.InitConf.CattleJoinedType {
		for _, v := range initData["cattleJoinedTypes"] {
			_, err := Client.CattleJoinedType.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("CattleJoinedType %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("CattleJoinedType %s initialized", v))
		}
		conf.InitConf.CattleJoinedType = true
	}

	// init CattleHairColor
	if !conf.InitConf.CattleHairColor {
		for _, v := range initData["cattleHairColors"] {
			_, err := Client.CattleHairColor.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("CattleHairColor %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("CattleHairColor %s initialized", v))
		}
		conf.InitConf.CattleHairColor = true
	}

	// init BreedingType
	if !conf.InitConf.BreedingType {
		for _, v := range initData["breedingTypes"] {
			_, err := Client.BreedingType.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("BreedingType %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("BreedingType %s initialized", v))
		}
		conf.InitConf.BreedingType = true
	}

	// init CattleOwner
	if !conf.InitConf.CattleOwner {
		for _, v := range initData["cattleOwners"] {
			_, err := Client.CattleOwner.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("CattleOwner %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("CattleOwner %s initialized", v))
		}
		conf.InitConf.CattleOwner = true
	}

	if !conf.InitConf.CalveType {
		for _, v := range initData["calveTypes"] {
			_, err := Client.CalveType.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("CalveType %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("CalveType %s initialized", v))
		}
		conf.InitConf.CalveType = true
	}

	if !conf.InitConf.CalveCount {
		for _, v := range initData["calveCounts"] {
			_, err := Client.CalveCount.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("CalveCount %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("CalveCount %s initialized", v))
		}
		conf.InitConf.CalveCount = true
	}

	if !conf.InitConf.EstrusType {
		for _, v := range initData["estrusTypes"] {
			_, err := Client.EstrusType.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("EstrusType %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("EstrusType %s initialized", v))
		}
		conf.InitConf.EstrusType = true
	}

	if !conf.InitConf.SemenFrozenType {
		for _, v := range initData["semenFrozenTypes"] {
			_, err := Client.SemenFrozenType.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("SemenFrozenType %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("SemenFrozenType %s initialized", v))
		}
		conf.InitConf.SemenFrozenType = true
	}

	if !conf.InitConf.PregnancyTestType {
		for _, v := range initData["pregnancyTestTypes"] {
			_, err := Client.PregnancyTestType.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("PregnancyTestType %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("PregnancyTestType %s initialized", v))
		}
		conf.InitConf.PregnancyTestType = true
	}

	if !conf.InitConf.PregnancyTestMethod {
		for _, v := range initData["pregnancyTestMethods"] {
			_, err := Client.PregnancyTestMethod.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("PregnancyTestMethod %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("PregnancyTestMethod %s initialized", v))
		}
		conf.InitConf.PregnancyTestMethod = true
	}

	if !conf.InitConf.PregnancyTestResult {
		for _, v := range initData["pregnancyTestResults"] {
			_, err := Client.PregnancyTestResult.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("PregnancyTestResult %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("PregnancyTestResult %s initialized", v))
		}
		conf.InitConf.PregnancyTestResult = true
	}

	if !conf.InitConf.AbortionType {
		for _, v := range initData["abortionTypes"] {
			_, err := Client.AbortionType.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("AbortionType %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("AbortionType %s initialized", v))
		}
		conf.InitConf.AbortionType = true
	}

	if !conf.InitConf.TreatmentResult {
		for _, v := range initData["treatmentResults"] {
			_, err := Client.TreatmentResult.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("TreatmentResult %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("TreatmentResult %s initialized", v))
		}
		conf.InitConf.TreatmentResult = true
	}

	if !conf.InitConf.TreatmentState {
		for _, v := range initData["treatmentStates"] {
			_, err := Client.TreatmentState.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("TreatmentState %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("TreatmentState %s initialized", v))
		}
		conf.InitConf.TreatmentState = true
	}

	if !conf.InitConf.Whereablout {
		for _, v := range initData["whereablouts"] {
			_, err := Client.Whereabouts.Create().SetName(v).Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("Whereablout %s init failed with error: %s", v, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("Whereablout %s initialized", v))
		}
		conf.InitConf.Whereablout = true
	}

	if !conf.InitConf.Api {
		for _, v := range apis {
			_, err := Client.API.Create().SetName(v.Name).SetID(v.ID).
				SetLevel(v.Level).SetComponent(v.Component).SetCreatedAt(v.CreatedAt).
				SetDeleted(v.Deleted).SetHasSub(v.HasSub).SetHash(v.Hash).SetIsSub(v.IsSub).
				SetParentId(v.ParentId).SetPath(v.Path).SetRedirect(v.Redirect).SetSingle(v.Single).
				SetTenantId(v.TenantId).SetUpdatedAt(v.UpdatedAt).
				Save(ctx)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("Api %s init failed with error: %s", v.Name, err.Error()))
				continue
			}
			log.Info().Msg(fmt.Sprintf("Api %s initialized", v.Name))
		}
		conf.InitConf.Api = true
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
