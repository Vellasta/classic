package encounters

import (
	"github.com/Vellasta/classic/sim/core"
	"github.com/Vellasta/classic/sim/core/proto"
	"github.com/Vellasta/classic/sim/core/stats"
)

func addPatchwerk(bossPrefix string) {
	core.AddPresetTarget(&core.PresetTarget{
		PathPrefix: bossPrefix,
		Config: &proto.Target{
			Id:        16028,
			Name:      "Patchwerk",
			Level:     63,
			MobType:   proto.MobType_MobTypeUndead,
			TankIndex: 0,

			Stats: stats.Stats{
				stats.Health:      3997200, // TODO:
				stats.Armor:       4611,    // TODO:
				stats.AttackPower: 805,     // TODO:
			}.ToFloatArray(),

			SpellSchool:      proto.SpellSchool_SpellSchoolPhysical,
			SwingSpeed:       2,      // TODO:
			MinBaseDamage:    3000,   // TODO:
			DamageSpread:     0.3333, // TODO:
			ParryHaste:       true,
			DualWield:        false,
			DualWieldPenalty: false,
			TargetInputs:     make([]*proto.TargetInput, 0),
		},
	})
	core.AddPresetEncounter("Patchwerk", []string{
		bossPrefix + "/Patchwerk",
	})
}

func addKelthuzad(bossPrefix string) {
	core.AddPresetTarget(&core.PresetTarget{
		PathPrefix: bossPrefix,
		Config: &proto.Target{
			Id:        15990,
			Name:      "Kel'Thuzad",
			Level:     63,
			MobType:   proto.MobType_MobTypeUndead,
			TankIndex: 0,

			Stats: stats.Stats{
				stats.Health:      3198000, // TODO:
				stats.Armor:       3402,    // TODO:
				stats.AttackPower: 805,     // TODO:
			}.ToFloatArray(),

			SpellSchool:      proto.SpellSchool_SpellSchoolPhysical,
			SwingSpeed:       2,      // TODO:
			MinBaseDamage:    3000,   // TODO:
			DamageSpread:     0.3333, // TODO:
			ParryHaste:       true,
			DualWield:        false,
			DualWieldPenalty: false,
			TargetInputs:     make([]*proto.TargetInput, 0),
		},
	})
	core.AddPresetEncounter("Kel'Thuzad", []string{
		bossPrefix + "/Kel'Thuzad",
	})
}

func addKeeperGnarlmoon(bossPrefix string) {
	core.AddPresetTarget(&core.PresetTarget{
		PathPrefix: bossPrefix,
		Config: &proto.Target{
			Id:        61939,
			Name:      "Keeper Gnarlmoon",
			Level:     63,
			MobType:   proto.MobType_MobTypeDemon,
			TankIndex: 0,

			Stats: stats.Stats{
				stats.Health:      3739000, // TODO:
				stats.Armor:       4211,    // TODO:
				stats.AttackPower: 805,     // TODO:
			}.ToFloatArray(),

			SpellSchool:      proto.SpellSchool_SpellSchoolPhysical,
			SwingSpeed:       2,      // TODO:
			MinBaseDamage:    3000,   // TODO:
			DamageSpread:     0.3333, // TODO:
			ParryHaste:       true,
			DualWield:        false,
			DualWieldPenalty: false,
			TargetInputs:     make([]*proto.TargetInput, 0),
		},
	})
	core.AddPresetEncounter("Keeper Gnarlmoon", []string{
		bossPrefix + "/Keeper Gnarlmoon",
	})
}

func addKruul(bossPrefix string) {
	core.AddPresetTarget(&core.PresetTarget{
		PathPrefix: bossPrefix,
		Config: &proto.Target{
			Id:        59991,
			Name:      "Kruul",
			Level:     63,
			MobType:   proto.MobType_MobTypeDemon,
			TankIndex: 0,

			Stats: stats.Stats{
				stats.Health:      4700000, // TODO:
				stats.Armor:       4752,    // TODO:
				stats.AttackPower: 805,     // TODO:
			}.ToFloatArray(),

			SpellSchool:      proto.SpellSchool_SpellSchoolPhysical,
			SwingSpeed:       2,      // TODO:
			MinBaseDamage:    3000,   // TODO:
			DamageSpread:     0.3333, // TODO:
			ParryHaste:       true,
			DualWield:        false,
			DualWieldPenalty: false,
			TargetInputs:     make([]*proto.TargetInput, 0),
		},
	})
	core.AddPresetEncounter("Kruul", []string{
		bossPrefix + "/Kruul",
	})
}

func addNothThePlaguebringer(bossPrefix string) {
	core.AddPresetTarget(&core.PresetTarget{
		PathPrefix: bossPrefix,
		Config: &proto.Target{
			Id:        15954,
			Name:      "Noth the Plaguebringer",
			Level:     63,
			MobType:   proto.MobType_MobTypeUndead,
			TankIndex: 0,

			Stats: stats.Stats{
				stats.Health:      1665500, // TODO:
				stats.Armor:       3850,    // TODO:
				stats.AttackPower: 805,     // TODO:
			}.ToFloatArray(),

			SpellSchool:      proto.SpellSchool_SpellSchoolPhysical,
			SwingSpeed:       2,      // TODO:
			MinBaseDamage:    3000,   // TODO:
			DamageSpread:     0.3333, // TODO:
			ParryHaste:       true,
			DualWield:        false,
			DualWieldPenalty: false,
			TargetInputs:     make([]*proto.TargetInput, 0),
		},
	})
	core.AddPresetEncounter("Noth the Plaguebringer", []string{
		bossPrefix + "/Noth the Plaguebringer",
	})
}

// func addLevel60(bossPrefix string) {
// 	core.AddPresetTarget(&core.PresetTarget{
// 		PathPrefix: bossPrefix,
// 		Config: &proto.Target{
// 			Id:        213336, // TODO:
// 			Name:      "Level 60",
// 			Level:     60,
// 			MobType:   proto.MobType_MobTypeUnknown,
// 			TankIndex: 0,

// 			Stats: stats.Stats{
// 				stats.Health:      127_393, // TODO:
// 				stats.Armor:       3731,    // TODO:
// 				stats.AttackPower: 805,     // TODO:
// 			}.ToFloatArray(),

// 			SpellSchool:      proto.SpellSchool_SpellSchoolPhysical,
// 			SwingSpeed:       2,      // TODO:
// 			MinBaseDamage:    3000,   // TODO:
// 			DamageSpread:     0.3333, // TODO:
// 			ParryHaste:       true,
// 			DualWield:        false,
// 			DualWieldPenalty: false,
// 			TargetInputs:     make([]*proto.TargetInput, 0),
// 		},
// 	})
// 	core.AddPresetEncounter("Level 60", []string{
// 		bossPrefix + "/Level 60",
// 	})
// }
