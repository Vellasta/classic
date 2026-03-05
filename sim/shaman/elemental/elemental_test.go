package elemental

import (
	"testing"

	_ "github.com/Vellasta/classic/sim/common"
	"github.com/Vellasta/classic/sim/core"
	"github.com/Vellasta/classic/sim/core/proto"
)

func init() {
	RegisterElementalShaman()
}

func TestElemental(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassShaman,
			Phase:      1,
			Race:       proto.Race_RaceTroll,
			OtherRaces: []proto.Race{proto.Race_RaceOrc},

			Talents:     DefaultTalents,
			GearSet:     core.GetGearSet("../../../ui/elemental_shaman/gear_sets", "phase_1"),
			Rotation:    core.GetAplRotation("../../../ui/elemental_shaman/apls", "default"),
			Buffs:       core.FullBuffs,
			Consumes:    Phase1Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Adaptive", SpecOptions: PlayerOptionsAdaptive},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,
		},
		{
			Class:      proto.Class_ClassShaman,
			Phase:      2,
			Race:       proto.Race_RaceTroll,
			OtherRaces: []proto.Race{proto.Race_RaceOrc},

			Talents:     DefaultTalents,
			GearSet:     core.GetGearSet("../../../ui/elemental_shaman/gear_sets", "phase_2"),
			Rotation:    core.GetAplRotation("../../../ui/elemental_shaman/apls", "default"),
			Buffs:       core.FullBuffs,
			Consumes:    Phase2Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Adaptive", SpecOptions: PlayerOptionsAdaptive},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,
		},
		{
			Class:      proto.Class_ClassShaman,
			Phase:      3,
			Race:       proto.Race_RaceTroll,
			OtherRaces: []proto.Race{proto.Race_RaceOrc},

			Talents:     DefaultTalents,
			GearSet:     core.GetGearSet("../../../ui/elemental_shaman/gear_sets", "phase_3"),
			Rotation:    core.GetAplRotation("../../../ui/elemental_shaman/apls", "default"),
			Buffs:       core.FullBuffs,
			Consumes:    Phase2Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Adaptive", SpecOptions: PlayerOptionsAdaptive},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,
		},
		{
			Class:      proto.Class_ClassShaman,
			Phase:      4,
			Race:       proto.Race_RaceTroll,
			OtherRaces: []proto.Race{proto.Race_RaceOrc},

			Talents:     DefaultTalents,
			GearSet:     core.GetGearSet("../../../ui/elemental_shaman/gear_sets", "phase_4"),
			Rotation:    core.GetAplRotation("../../../ui/elemental_shaman/apls", "default"),
			Buffs:       core.FullBuffs,
			Consumes:    Phase2Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Adaptive", SpecOptions: PlayerOptionsAdaptive},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,
		},
		{
			Class:      proto.Class_ClassShaman,
			Phase:      5,
			Race:       proto.Race_RaceTroll,
			OtherRaces: []proto.Race{proto.Race_RaceOrc},

			Talents:     DefaultTalents,
			GearSet:     core.GetGearSet("../../../ui/elemental_shaman/gear_sets", "phase_5"),
			Rotation:    core.GetAplRotation("../../../ui/elemental_shaman/apls", "default"),
			Buffs:       core.FullBuffs,
			Consumes:    Phase5Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Adaptive", SpecOptions: PlayerOptionsAdaptive},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,
		},
		{
			Class:      proto.Class_ClassShaman,
			Phase:      6,
			Race:       proto.Race_RaceTroll,
			OtherRaces: []proto.Race{proto.Race_RaceOrc},

			Talents:     DefaultTalents,
			GearSet:     core.GetGearSet("../../../ui/elemental_shaman/gear_sets", "phase_6"),
			Rotation:    core.GetAplRotation("../../../ui/elemental_shaman/apls", "default"),
			Buffs:       core.FullBuffs,
			Consumes:    Phase5Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Adaptive", SpecOptions: PlayerOptionsAdaptive},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,
		},
	}))
}

var DefaultTalents = "550331050002151--50105301005"

var PlayerOptionsAdaptive = &proto.Player_ElementalShaman{
	ElementalShaman: &proto.ElementalShaman{
		Options: &proto.ElementalShaman_Options{},
	},
}

var Phase1Consumes = core.ConsumesCombo{
	Label: "P1-Consumes",
	Consumes: &proto.Consumes{
		DefaultConjured: proto.Conjured_ConjuredDemonicRune,
		DefaultPotion:   proto.Potions_MajorManaPotion,
		Flask:           proto.Flask_FlaskOfSupremePower,
		FirePowerBuff:   proto.FirePowerBuff_ElixirOfGreaterFirepower,
		Food:            proto.Food_FoodNightfinSoup,
		MainHandImbue:   proto.WeaponImbue_LesserWizardOil,
		SpellPowerBuff:  proto.SpellPowerBuff_GreaterArcaneElixir,
	},
}

var Phase2Consumes = core.ConsumesCombo{
	Label: "P2-Consumes",
	Consumes: &proto.Consumes{
		DefaultConjured: proto.Conjured_ConjuredDemonicRune,
		DefaultPotion:   proto.Potions_MajorManaPotion,
		Flask:           proto.Flask_FlaskOfSupremePower,
		FirePowerBuff:   proto.FirePowerBuff_ElixirOfGreaterFirepower,
		Food:            proto.Food_FoodRunnTumTuberSurprise,
		MainHandImbue:   proto.WeaponImbue_LesserWizardOil,
		SpellPowerBuff:  proto.SpellPowerBuff_GreaterArcaneElixir,
	},
}

var Phase5Consumes = core.ConsumesCombo{
	Label: "P5-Consumes",
	Consumes: &proto.Consumes{
		DefaultConjured: proto.Conjured_ConjuredDemonicRune,
		DefaultPotion:   proto.Potions_MajorManaPotion,
		Flask:           proto.Flask_FlaskOfSupremePower,
		FirePowerBuff:   proto.FirePowerBuff_ElixirOfGreaterFirepower,
		Food:            proto.Food_FoodRunnTumTuberSurprise,
		MainHandImbue:   proto.WeaponImbue_BrilliantWizardOil,
		SpellPowerBuff:  proto.SpellPowerBuff_GreaterArcaneElixir,
	},
}

var ItemFilters = core.ItemFilter{
	WeaponTypes: []proto.WeaponType{
		proto.WeaponType_WeaponTypeAxe,
		proto.WeaponType_WeaponTypeDagger,
		proto.WeaponType_WeaponTypeFist,
		proto.WeaponType_WeaponTypeMace,
		proto.WeaponType_WeaponTypeOffHand,
		proto.WeaponType_WeaponTypeShield,
		proto.WeaponType_WeaponTypeStaff,
	},
	ArmorType: proto.ArmorType_ArmorTypeMail,
	RangedWeaponTypes: []proto.RangedWeaponType{
		proto.RangedWeaponType_RangedWeaponTypeTotem,
	},
}

var Stats = []proto.Stat{
	proto.Stat_StatIntellect,
	proto.Stat_StatSpellPower,
	proto.Stat_StatNaturePower,
	proto.Stat_StatFirePower,
	proto.Stat_StatSpellHit,
	proto.Stat_StatSpellCrit,
}
