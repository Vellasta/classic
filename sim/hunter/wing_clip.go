package hunter

import (
	"time"

	"github.com/Vellasta/classic/sim/core"
)

func (hunter *Hunter) getWingClipConfig(rank int) core.SpellConfig {
	spellId := [4]int32{0, 2974, 14267, 14268}[rank]
	wingClipWeaponDamage := [4]float64{0, 0.25, 0.3, 0.35}[rank]
	manaCost := [4]float64{0, 40, 60, 80}[rank]
	level := [4]int{0, 12, 38, 60}[rank]

	return core.SpellConfig{
		SpellCode:     SpellCode_HunterWingClip,
		ActionID:      core.ActionID{SpellID: spellId},
		SpellSchool:   core.SpellSchoolPhysical,
		DefenseType:   core.DefenseTypeMelee,
		ProcMask:      core.ProcMaskMeleeMHSpecial,
		Flags:         core.SpellFlagMeleeMetrics | core.SpellFlagAPL | core.SpellFlagBinary,
		Rank:          rank,
		RequiredLevel: level,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost * (1 - 0.02*float64(hunter.Talents.Resourcefulness)),
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    hunter.NewTimer(),
				Duration: time.Millisecond * 3000,
			},
			IgnoreHaste: true,
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return hunter.DistanceFromTarget <= core.MaxMeleeAttackDistance
		},

		BonusCritRating:  float64(hunter.Talents.SavageStrikes) * 3 * core.CritRatingPerCritChance,
		CritDamageBonus:  0.066 * float64(hunter.Talents.KillerInstinct),
		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			damageMH := wingClipWeaponDamage * hunter.MHWeaponDamage(sim, spell.MeleeAttackPower(target))
			spell.CalcAndDealDamage(sim, target, damageMH, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
		},
	}
}

func (hunter *Hunter) registerWingClipSpell() {
	rank := map[int32]int{
		25: 1,
		40: 2,
		50: 3,
		60: 3,
	}[hunter.Level]

	config := hunter.getWingClipConfig(rank)
	hunter.WingClip = hunter.GetOrRegisterSpell(config)
}
