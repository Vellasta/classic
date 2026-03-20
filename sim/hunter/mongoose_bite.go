package hunter

import (
	"time"

	"github.com/Vellasta/classic/sim/core"
)

func (hunter *Hunter) getMongooseBiteConfig(rank int) core.SpellConfig {
	spellId := [5]int32{0, 1495, 14269, 14270, 14271}[rank]
	mongooseBiteBaseDamage := [5]float64{0, 15, 25, 45, 70}[rank]
	mongooseBiteWeaponDamage := [5]float64{0, 0.45, 0.5, 0.55, 0.6}[rank]
	viciousStrikesCDReduction := [3]time.Duration{0, 500 * time.Millisecond, 1000 * time.Millisecond}[hunter.Talents.ViciousStrikes]
	manaCost := [5]float64{0, 30, 40, 50, 65}[rank]
	level := [5]int{0, 16, 30, 44, 58}[rank]

	spellConfig := core.SpellConfig{
		SpellCode:     SpellCode_HunterMongooseBite,
		ActionID:      core.ActionID{SpellID: spellId},
		SpellSchool:   core.SpellSchoolPhysical,
		DefenseType:   core.DefenseTypeMelee,
		ProcMask:      core.ProcMaskMeleeSpecial,
		Flags:         core.SpellFlagMeleeMetrics | core.SpellFlagAPL,
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
				Duration: time.Millisecond*5000 - viciousStrikesCDReduction,
			},
		},

		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return hunter.DistanceFromTarget <= core.MaxMeleeAttackDistance
		},

		BonusCritRating:  float64(hunter.Talents.SavageStrikes) * 3 * core.CritRatingPerCritChance,
		CritDamageBonus:  0.066 * float64(hunter.Talents.KillerInstinct),
		DamageMultiplier: 1 + 0.05*float64(hunter.Talents.ViciousStrikes),
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			// hunter.DefensiveState.Deactivate(sim)
			damageMH := mongooseBiteBaseDamage + mongooseBiteWeaponDamage*hunter.MHWeaponDamage(sim, spell.MeleeAttackPower(target))
			result := spell.CalcDamage(sim, target, damageMH, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
			spell.DealDamage(sim, result)
			if result.Landed() && (hunter.Talents.StingingNettle > 0) {
				dot := hunter.SerpentSting.Dot(target)
				dot.NumberOfTicks = hunter.Talents.StingingNettle
				dot.Apply(sim)
			}
			if hunter.HasOHWeapon() {
				damageOH := mongooseBiteBaseDamage + mongooseBiteWeaponDamage*hunter.OHWeaponDamage(sim, spell.MeleeAttackPower(target))
				spell.CalcAndDealDamage(sim, target, damageOH, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
			}
		},
	}

	return spellConfig
}

func (hunter *Hunter) registerMongooseBiteSpell() {
	// Aura is only used as a pre-requisite for Mongoose Bite
	// hunter.DefensiveState = hunter.RegisterAura(core.Aura{
	// 	Label:    "Defensive State",
	// 	ActionID: core.ActionID{SpellID: 5302},
	// 	Duration: time.Second * 5,

	// 	OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
	// 		if result.DidDodge() {
	// 			aura.Activate(sim)
	// 		}
	// 	},
	// })

	rank := map[int32]int{
		25: 1,
		40: 2,
		50: 3,
		60: 4,
	}[hunter.Level]

	config := hunter.getMongooseBiteConfig(rank)
	hunter.MongooseBite = hunter.GetOrRegisterSpell(config)
}
