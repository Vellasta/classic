package warlock

import (
	"time"

	"github.com/Vellasta/classic/sim/core"
)

func (warlock *Warlock) registerSummonDemon() {
	manaCost := core.ManaCostOptions{
		FlatCost: warlock.BaseMana,
	}
	// All have a default cast time of 10s and the active pet is dismissed when the cast starts
	cast := core.CastConfig{
		DefaultCast: core.Cast{
			GCD:      core.GCDDefault,
			CastTime: time.Second * 10,
		},
		ModifyCast: func(sim *core.Simulation, spell *core.Spell, cast *core.Cast) {
			warlock.changeActivePet(sim, nil, false)
		},
	}

	// Felhunter
	warlock.SummonDemonSpells = append(warlock.SummonDemonSpells, warlock.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 691},
		SpellSchool: core.SpellSchoolShadow,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagAPL,

		ManaCost: manaCost,
		Cast:     cast,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			warlock.changeActivePet(sim, warlock.Felhunter, false)
		},
	}))

	// Imp
	warlock.SummonDemonSpells = append(warlock.SummonDemonSpells, warlock.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 688},
		SpellSchool: core.SpellSchoolShadow,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagAPL,

		ManaCost: manaCost,
		Cast:     cast,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			warlock.changeActivePet(sim, warlock.Imp, false)
		},
	}))

	// Succubus
	warlock.SummonDemonSpells = append(warlock.SummonDemonSpells, warlock.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 712},
		SpellSchool: core.SpellSchoolShadow,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagAPL,

		ManaCost: manaCost,
		Cast:     cast,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			warlock.changeActivePet(sim, warlock.Succubus, false)
		},
	}))

	// Voidwalker
	warlock.SummonDemonSpells = append(warlock.SummonDemonSpells, warlock.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 697},
		SpellSchool: core.SpellSchoolShadow,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagAPL,

		ManaCost: manaCost,
		Cast:     cast,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			warlock.changeActivePet(sim, warlock.Voidwalker, false)
		},
	}))
}
