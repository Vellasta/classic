package hunter

import (
	"strconv"
	"time"

	"github.com/Vellasta/classic/sim/core"
	"github.com/Vellasta/classic/sim/core/stats"
)

// Utility function to create an Improved Wolf Aura
func (hunter *Hunter) createImprovedWolfAura(auraLabel string, actionID core.ActionID) *core.Aura {
	return hunter.GetOrRegisterAura(core.Aura{
		Label:    auraLabel,
		ActionID: actionID,
		Duration: time.Second * 12,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.MultiplyMeleeSpeed(sim, hunter.SwiftAspectsMultiplier)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.MultiplyMeleeSpeed(sim, 1/hunter.SwiftAspectsMultiplier)
		},
	})
}

// Function to get the maximum attack power for Aspect of the Wolf based on rank
func (hunter *Hunter) getMaxAspectOfTheWolfAttackPower(rank int) float64 {
	attackPower := [8]float64{0, 20, 35, 50, 70, 90, 110, 120}

	if rank < 1 || rank > 7 {
		return 0.0
	}

	return attackPower[rank]
}

func (hunter *Hunter) getMaxWolfRank() int {
	maxRank := core.TernaryInt(core.IncludeAQ, 7, 6)

	for i := maxRank; i > 0; i-- {
		config := hunter.getAspectOfTheWolfSpellConfig(i)
		if config.RequiredLevel <= int(hunter.Level) {
			return i
		}
	}
	return 1
}

func (hunter *Hunter) getAspectOfTheWolfSpellConfig(rank int) core.SpellConfig {
	improvedWolfProcChance := 0.1

	spellIds := [8]int32{0, 45650, 51496, 51497, 51498, 51499, 51500, 51501}
	levels := [8]int{0, 10, 18, 28, 38, 48, 58, 60}

	spellId := spellIds[rank]
	level := levels[rank]

	if hunter.Talents.SwiftAspects > 0 {
		hunter.impWolfAura = hunter.createImprovedWolfAura(
			"Quick Strikes",
			core.ActionID{SpellID: 51546},
		)
	}
	// Use utility function to get the attack power based on rank
	meleeap := hunter.getMaxAspectOfTheWolfAttackPower(rank)

	actionID := core.ActionID{SpellID: spellId}
	aspectOfTheWolfAura := hunter.GetOrRegisterAura(core.Aura{
		Label:    "Aspect of the Wolf" + strconv.Itoa(rank),
		ActionID: actionID,
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.AddStatDynamic(sim, stats.AttackPower, meleeap*hunter.AspectOfTheWolfAPMultiplier)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.AddStatDynamic(sim, stats.AttackPower, -meleeap*hunter.AspectOfTheWolfAPMultiplier)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.ProcMask.Matches(core.ProcMaskEmpty) || !spell.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}

			if hunter.impWolfAura != nil && sim.Proc(improvedWolfProcChance, "Imp Aspect of the Wolf") {
				hunter.impWolfAura.Activate(sim)
			}
		},
	})

	aspectOfTheWolfAura.NewExclusiveEffect("Aspect", true, core.ExclusiveEffect{})

	return core.SpellConfig{
		ActionID:      actionID,
		Flags:         core.SpellFlagAPL,
		Rank:          rank,
		RequiredLevel: level,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return !aspectOfTheWolfAura.IsActive()
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			aspectOfTheWolfAura.Activate(sim)
		},
	}
}

func (hunter *Hunter) registerAspectOfTheWolfSpell() {
	hunter.AspectOfTheWolfAPMultiplier = 1.0
	hunter.SwiftAspectsMultiplier = 1 + 0.03*float64(hunter.Talents.SwiftAspects)
	maxRank := hunter.getMaxWolfRank()
	config := hunter.getAspectOfTheWolfSpellConfig(maxRank)
	hunter.GetOrRegisterSpell(config)
}
