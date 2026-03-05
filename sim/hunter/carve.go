package hunter

import (
	"time"

	"github.com/Vellasta/classic/sim/core"
)

func (hunter *Hunter) getCarveConfig(rank int) core.SpellConfig {
	spellId := [4]int32{0, 51575, 52415, 52416}[rank]
	carveWeaponDamage := [4]float64{0, 0.6, 0.65, 0.7}[rank]
	level := [4]int{0, 20, 38, 56}[rank]
	numHits := hunter.Env.GetNumTargets()

	spellConfig := core.SpellConfig{
		SpellCode:     SpellCode_HunterCarve,
		ActionID:      core.ActionID{SpellID: spellId},
		SpellSchool:   core.SpellSchoolPhysical,
		DefenseType:   core.DefenseTypeMelee,
		ProcMask:      core.ProcMaskMeleeSpecial,
		Flags:         core.SpellFlagMeleeMetrics | core.SpellFlagAPL,
		Rank:          rank,
		RequiredLevel: level,

		ManaCost: core.ManaCostOptions{
			FlatCost: 0.08 * hunter.BaseMana * (1 - 0.02*float64(hunter.Talents.Resourcefulness)),
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    hunter.NewTimer(),
				Duration: time.Second * 10,
			},
		},

		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return hunter.DistanceFromTarget <= core.MaxMeleeAttackDistance
		},

		BonusCritRating:  float64(hunter.Talents.SavageStrikes) * 3 * core.CritRatingPerCritChance,
		CritDamageBonus:  0.066 * float64(hunter.Talents.KillerInstinct),
		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			curTarget := target
			for hitIndex := int32(0); hitIndex < numHits; hitIndex++ {
				baseDamage := carveWeaponDamage * hunter.MHWeaponDamage(sim, spell.MeleeAttackPower(target))
				spell.CalcAndDealDamage(sim, curTarget, baseDamage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
				curTarget = sim.Environment.NextTargetUnit(curTarget)
			}
		},
	}

	return spellConfig
}

func (hunter *Hunter) registerCarveSpell() {
	if !hunter.Talents.Carve {
		return
	}

	rank := map[int32]int{
		25: 1,
		40: 2,
		50: 3,
		60: 3,
	}[hunter.Level]

	config := hunter.getCarveConfig(rank)
	hunter.Carve = hunter.GetOrRegisterSpell(config)
}
