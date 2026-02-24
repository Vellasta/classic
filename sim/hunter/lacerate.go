package hunter

import (
	"strconv"
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (hunter *Hunter) getLacerateConfig(rank int) core.SpellConfig {
	spellId := [2]int32{0, 48049}[rank]
	lacerateWeaponDamage := [2]float64{0, 0.40}[rank]
	level := [2]int{0, 40}[rank]

	spellConfig := core.SpellConfig{
		SpellCode:     SpellCode_HunterLacerate,
		ActionID:      core.ActionID{SpellID: spellId},
		SpellSchool:   core.SpellSchoolPhysical,
		DefenseType:   core.DefenseTypeMelee,
		ProcMask:      core.ProcMaskMeleeSpecial,
		Flags:         core.SpellFlagMeleeMetrics | core.SpellFlagAPL,
		Rank:          rank,
		RequiredLevel: level,

		ManaCost: core.ManaCostOptions{
			FlatCost: 0.08 * hunter.BaseMana,
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
			return hunter.DistanceFromTarget <= core.MaxMeleeAttackDistance && hunter.LacerateAura.IsActive()
		},

		BonusCritRating:  core.CritRatingPerCritChance,
		CritDamageBonus:  0.066 * float64(hunter.Talents.KillerInstinct),
		DamageMultiplier: 1.15,
		ThreatMultiplier: 1,

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: "Lacerate" + hunter.Label + strconv.Itoa(rank),
				Tag:   "Lacerate",
			},
			NumberOfTicks: 4,
			TickLength:    time.Millisecond * 2000,

			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := lacerateWeaponDamage * spell.MeleeAttackPower(target)
			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
			spell.DealDamage(sim, result)
			if result.Landed() {
				dot := spell.Dot(target)
				tickDamage := (result.Damage * 0.2) / float64(dot.NumberOfTicks)
				dot.Snapshot(target, tickDamage, false)
				dot.Apply(sim)
			}
		},
	}

	return spellConfig
}

func (hunter *Hunter) registerLacerateSpell() {
	if !hunter.Talents.Lacerate {
		return
	}

	hunter.RegisterAura(core.Aura{
		Label:    "Lacerate Trigger",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if result.DidCrit() {
				hunter.LacerateAura.Activate(sim)
			}
		},
	})

	hunter.LacerateAura = hunter.RegisterAura(core.Aura{
		Label:    "Lacerate Aura",
		ActionID: core.ActionID{SpellID: 48049},
		Duration: time.Second * 5,
	})

	rank := map[int32]int{
		25: 1,
		40: 1,
		50: 1,
		60: 1,
	}[hunter.Level]

	config := hunter.getLacerateConfig(rank)
	hunter.Lacerate = hunter.GetOrRegisterSpell(config)
}
