package paladin

import (
	"time"

	"github.com/Vellasta/classic/sim/core"
)

func (paladin *Paladin) registerHolyShock() {
	if !paladin.Talents.HolyShock {
		return
	}

	ranks := []struct {
		level     int32
		spellID   int32
		manaCost  float64
		minDamage float64
		maxDamage float64
	}{
		{level: 40, spellID: 20473, manaCost: 225, minDamage: 204, maxDamage: 220},
		{level: 48, spellID: 20929, manaCost: 275, minDamage: 279, maxDamage: 301},
		{level: 56, spellID: 20930, manaCost: 325, minDamage: 365, maxDamage: 395},
	}

	for i, rank := range ranks {
		rank := rank
		if paladin.Level < rank.level {
			break
		}

		paladin.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: rank.spellID},
			SpellSchool: core.SpellSchoolHoly,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellDamage,
			Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagAPL,

			RequiredLevel: int(rank.level),
			Rank:          i + 1,

			SpellCode: SpellCode_PaladinHolyShock,

			ManaCost: core.ManaCostOptions{
				FlatCost: rank.manaCost,
			},

			Cast: core.CastConfig{
				DefaultCast: core.Cast{
					GCD: core.GCDDefault,
				},
				CD: core.Cooldown{
					Timer:    paladin.NewTimer(),
					Duration: time.Second * 30,
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BonusCoefficient: 0.429,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				baseDamage := sim.Roll(rank.minDamage, rank.maxDamage)
				spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
			},
		})
	}
}
