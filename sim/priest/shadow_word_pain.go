package priest

import (
	"fmt"
	"time"

	"github.com/Vellasta/classic/sim/core"
)

const ShadowWordPainRanks = 8

var ShadowWordPainSpellId = [ShadowWordPainRanks + 1]int32{0, 589, 594, 970, 992, 2767, 10892, 10893, 10894}
var ShadowWordPainBaseDamage = [ShadowWordPainRanks + 1]float64{0, 30, 66, 132, 234, 366, 510, 672, 852}
var ShadowWordPainSpellCoef = [ShadowWordPainRanks + 1]float64{0, 0.067, 0.104, 0.154, 0.167, 0.167, 0.167, 0.167, 0.167} // per tick
var ShadowWordPainManaCost = [ShadowWordPainRanks + 1]float64{0, 25, 50, 95, 155, 230, 305, 385, 470}
var ShadowWordPainLevel = [ShadowWordPainRanks + 1]int{0, 4, 10, 18, 26, 34, 42, 50, 58}

//To Do: Check rollover code from runes

func (priest *Priest) registerShadowWordPainSpell() {
	priest.ShadowWordPain = make([]*core.Spell, ShadowWordPainRanks+1)

	for rank := 1; rank <= ShadowWordPainRanks; rank++ {
		config := priest.getShadowWordPainConfig(rank)

		if config.RequiredLevel <= int(priest.Level) {
			priest.ShadowWordPain[rank] = priest.GetOrRegisterSpell(config)
		}
	}
}

func (priest *Priest) getShadowWordPainConfig(rank int) core.SpellConfig {
	ticks := int32(6)

	spellId := ShadowWordPainSpellId[rank]
	baseDotDamage := ShadowWordPainBaseDamage[rank] / float64(ticks)
	spellCoeff := ShadowWordPainSpellCoef[rank]
	manaCost := ShadowWordPainManaCost[rank]
	level := ShadowWordPainLevel[rank]

	return core.SpellConfig{
		SpellCode:   SpellCode_PriestShadowWordPain,
		ActionID:    core.ActionID{SpellID: spellId},
		SpellSchool: core.SpellSchoolShadow,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       SpellFlagPriest | core.SpellFlagAPL | core.SpellFlagPureDot,

		RequiredLevel: level,
		Rank:          rank,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: fmt.Sprintf("Shadow Word: Pain (Rank %d)", rank),
			},

			NumberOfTicks:    ticks + (priest.Talents.ImprovedShadowWordPain),
			TickLength:       time.Second * 3,
			BonusCoefficient: spellCoeff,

			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
				dot.Snapshot(target, baseDotDamage, isRollover)
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcOutcome(sim, target, spell.OutcomeMagicHitNoHitCounter)

			if result.Landed() {
				priest.AddShadowWeavingStack(sim, result.Target)
				spell.Dot(result.Target).Apply(sim)
			}
			spell.DealOutcome(sim, result)
		},
	}
}
