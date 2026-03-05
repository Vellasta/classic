package shaman

import (
	"fmt"
	"time"

	"github.com/Vellasta/classic/sim/core"
)

func (shaman *Shaman) setActiveAirTotem(sim *core.Simulation, spell *core.Spell, aura *core.Aura) {
	shaman.TotemExpirations[AirTotem] = sim.CurrentTime + aura.Duration
	shaman.ActiveTotems[AirTotem] = spell

	if shaman.ActiveTotemBuffs[AirTotem] != nil {
		shaman.ActiveTotemBuffs[AirTotem].Deactivate(sim)
	}

	shaman.ActiveTotemBuffs[AirTotem] = aura
	aura.Activate(sim)
}

const WindfuryTotemRanks = 3

var WindfuryTotemSpellId = [WindfuryTotemRanks + 1]int32{0, 8512, 10613, 10614}
var WindfuryBuffAuraId = [WindfuryTotemRanks + 1]int32{0, 8514, 10607, 10611}
var WindfuryTotemBonusDamage = [WindfuryTotemRanks + 1]float64{0, 122, 229, 315}
var WindfuryTotemManaCost = [WindfuryTotemRanks + 1]float64{0, 115, 175, 250}
var WindfuryTotemLevel = [WindfuryTotemRanks + 1]int{0, 32, 42, 52}

func (shaman *Shaman) registerWindfuryTotemSpell() {
	shaman.WindfuryTotem = make([]*core.Spell, WindfuryTotemRanks+1)
	shaman.WindfuryTotemPeriodicActions = make([]*core.PendingAction, WindfuryTotemRanks+1)

	for rank := 1; rank <= WindfuryTotemRanks; rank++ {
		config := shaman.newWindfuryTotemSpellConfig(rank)

		if config.RequiredLevel <= int(shaman.Level) {
			shaman.WindfuryTotem[rank] = shaman.RegisterSpell(config)
		}
	}

	shaman.AirTotems = append(
		shaman.AirTotems,
		core.FilterSlice(shaman.WindfuryTotem, func(spell *core.Spell) bool { return spell != nil })...,
	)
}

func (shaman *Shaman) newWindfuryTotemSpellConfig(rank int) core.SpellConfig {
	spellId := WindfuryTotemSpellId[rank]
	// TODO: The sim won't respect the value of a totem dropped via the APL. It uses hard-coded values from buffs.go
	// bonusDamage := WindfuryTotemBonusDamage[rank]
	manaCost := WindfuryTotemManaCost[rank]
	level := WindfuryTotemLevel[rank]

	// Create a trackable aura for totem weaving
	buffAura := shaman.RegisterAura(core.Aura{
		ActionID: core.ActionID{SpellID: WindfuryBuffAuraId[rank]},
		Label:    fmt.Sprintf("Windfury (Rank %d)", rank),
		Duration: time.Second * 10,
	})

	periodicTriggerAura := shaman.RegisterAura(core.Aura{
		Label:    fmt.Sprintf("Windfury Trigger Dummy (Rank %d)", rank),
		Duration: time.Minute * 2,
		OnGain: func(_ *core.Aura, sim *core.Simulation) {
			shaman.ActiveWindfuryTotemPeriodicAction = core.StartPeriodicAction(sim, core.PeriodicActionOptions{
				Period:          time.Second * 5, // Totem refreshes every 5 seconds
				TickImmediately: true,
				OnAction: func(_ *core.Simulation) {
					buffAura.Activate(sim)
				},
			})
		},
		OnExpire: func(_ *core.Aura, sim *core.Simulation) {
			shaman.ActiveWindfuryTotemPeriodicAction.Cancel(sim)
			shaman.ActiveWindfuryTotemPeriodicAction = nil
		},
	})

	spell := shaman.newTotemSpellConfig(manaCost, spellId)
	spell.RequiredLevel = level
	spell.Rank = rank
	spell.ApplyEffects = func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
		shaman.setActiveAirTotem(sim, spell, periodicTriggerAura)
	}
	return spell
}

const GraceOfAirTotemRanks = 3

var GraceOfAirTotemSpellId = [GraceOfAirTotemRanks + 1]int32{0, 8835, 10627, 25359}
var GraceOfAirTotemManaCost = [GraceOfAirTotemRanks + 1]float64{0, 155, 250, 310}
var GraceOfAirTotemLevel = [GraceOfAirTotemRanks + 1]int{0, 42, 56, 60}

func (shaman *Shaman) registerGraceOfAirTotemSpell() {
	shaman.GraceOfAirTotem = make([]*core.Spell, GraceOfAirTotemRanks+1)

	for rank := 1; rank <= GraceOfAirTotemRanks; rank++ {
		config := shaman.newGraceOfAirTotemSpellConfig(rank)

		if config.RequiredLevel <= int(shaman.Level) {
			shaman.GraceOfAirTotem[rank] = shaman.RegisterSpell(config)
		}
	}

	shaman.AirTotems = append(
		shaman.AirTotems,
		core.FilterSlice(shaman.GraceOfAirTotem, func(spell *core.Spell) bool { return spell != nil })...,
	)
}

func (shaman *Shaman) newGraceOfAirTotemSpellConfig(rank int) core.SpellConfig {
	spellId := GraceOfAirTotemSpellId[rank]
	manaCost := GraceOfAirTotemManaCost[rank]
	level := GraceOfAirTotemLevel[rank]

	multiplier := []float64{1, 1.08, 1.15}[shaman.Talents.EnhancingTotems]

	buffAura := core.GraceOfAirTotemAura(&shaman.Unit, multiplier)

	spell := shaman.newTotemSpellConfig(manaCost, spellId)
	spell.RequiredLevel = level
	spell.Rank = rank
	spell.ApplyEffects = func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
		shaman.setActiveAirTotem(sim, spell, buffAura)
	}
	return spell
}

const WindwallTotemRanks = 3

var WindwallTotemSpellId = [WindwallTotemRanks + 1]int32{0, 15107, 15111, 15112}
var WindwallTotemManaCost = [WindwallTotemRanks + 1]float64{0, 115, 170, 225}
var WindwallTotemLevel = [WindwallTotemRanks + 1]int{0, 36, 46, 56}

func (shaman *Shaman) registerWindwallTotemSpell() {
	shaman.WindwallTotem = make([]*core.Spell, WindwallTotemRanks+1)

	for rank := 1; rank <= WindwallTotemRanks; rank++ {
		config := shaman.newWindwallTotemSpellConfig(rank)

		if config.RequiredLevel <= int(shaman.Level) {
			shaman.WindwallTotem[rank] = shaman.RegisterSpell(config)
		}
	}

	shaman.AirTotems = append(
		shaman.AirTotems,
		core.FilterSlice(shaman.WindwallTotem, func(spell *core.Spell) bool { return spell != nil })...,
	)
}

func (shaman *Shaman) newWindwallTotemSpellConfig(rank int) core.SpellConfig {
	spellId := WindwallTotemSpellId[rank]
	manaCost := WindwallTotemManaCost[rank]
	level := WindwallTotemLevel[rank]

	spell := shaman.newTotemSpellConfig(manaCost, spellId)
	spell.RequiredLevel = level
	spell.Rank = rank
	spell.ApplyEffects = func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
		shaman.setActiveAirTotem(sim, spell, nil)
	}
	return spell
}
