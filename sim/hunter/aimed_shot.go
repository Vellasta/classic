package hunter

import (
	"time"

	"github.com/Vellasta/classic/sim/core"
	"github.com/Vellasta/classic/sim/core/proto"
)

func (hunter *Hunter) getAimedShotConfig(rank int, timer *core.Timer) core.SpellConfig {
	spellId := [7]int32{0, 19434, 20900, 20901, 20902, 20903, 20904}[rank]
	baseDamage := [7]float64{0, 70, 125, 200, 330, 460, 600}[rank]
	manaCost := [7]float64{0, 50, 75, 105, 140, 170, 205}[rank]
	level := [7]int{0, 0, 28, 36, 44, 52, 60}[rank]

	return core.SpellConfig{
		SpellCode:     SpellCode_HunterAimedShot,
		ActionID:      core.ActionID{SpellID: spellId},
		SpellSchool:   core.SpellSchoolPhysical,
		DefenseType:   core.DefenseTypeRanged,
		ProcMask:      core.ProcMaskRangedSpecial,
		Flags:         core.SpellFlagMeleeMetrics | core.SpellFlagAPL | SpellFlagShot,
		CastType:      proto.CastType_CastTypeRanged,
		Rank:          rank,
		RequiredLevel: level,
		MissileSpeed:  24,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:          core.GCDMin,
				BaseCastTime: time.Millisecond * 500,
				CastTime:     time.Millisecond * 2000,
			},
			CD: core.Cooldown{
				Timer:    timer,
				Duration: time.Second*26 - time.Millisecond*2000*time.Duration(hunter.Talents.Swiftshot),
			},
			ModifyCast: func(sim *core.Simulation, spell *core.Spell, cast *core.Cast) {
				cast.CastTime = spell.CastTime()
				hunter.Unit.AutoAttacks.CancelAutoSwing(sim)
			},
			IgnoreHaste: true, // Hunter GCD is locked at 1.5s
			CastTime: func(spell *core.Spell) time.Duration {
				return time.Duration(float64(spell.DefaultCast.BaseCastTime) + float64(spell.DefaultCast.CastTime)/hunter.RangedSwingSpeed())
			},
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return hunter.DistanceFromTarget >= core.MinRangedAttackDistance
		},

		CritDamageBonus: hunter.mortalShots(),

		DamageMultiplier: 1 + 0.05*float64(hunter.Talents.ImprovedMarksmanship),
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			curTarget := target
			numHits := int32(1)
			if hunter.lockAndLoadAura.IsActive() {
				numHits = hunter.Env.GetNumTargets()
			}
			results := make([]*core.SpellResult, numHits)

			for hitIndex := int32(0); hitIndex < numHits; hitIndex++ {
				baseDamage := hunter.AutoAttacks.Ranged().CalculateNormalizedWeaponDamage(sim, spell.RangedAttackPower(curTarget, false)) +
					hunter.NormalizedAmmoDamageBonus +
					baseDamage

				results[hitIndex] = spell.CalcDamage(sim, curTarget, baseDamage, spell.OutcomeRangedHitAndCrit)

				curTarget = sim.Environment.NextTargetUnit(curTarget)
			}
			hunter.Unit.AutoAttacks.EnableAutoSwing(sim)
			spell.WaitTravelTime(sim, func(s *core.Simulation) {
				if hunter.lockAndLoadAura.IsActive() {
					hunter.lockAndLoadAura.Deactivate(sim)
				}
				for hitIndex := int32(0); hitIndex < numHits; hitIndex++ {
					spell.DealDamage(sim, results[hitIndex])

					// Apply experimental ammo damage (calculated independently)
					ammo := hunter.ExperimentalAmmunitionTypes[hunter.ExperimentalAmmunitionState]
					ammo.Cast(sim, curTarget)

					curTarget = sim.Environment.NextTargetUnit(curTarget)
				}
			})
		},
	}
}

func (hunter *Hunter) registerAimedShotSpell(timer *core.Timer) {
	if !hunter.Talents.AimedShot {
		return
	}

	maxRank := 6

	for i := 1; i <= maxRank; i++ {
		config := hunter.getAimedShotConfig(i, timer)

		if config.RequiredLevel <= int(hunter.Level) {
			hunter.AimedShot = hunter.GetOrRegisterSpell(config)
		}
	}
}
