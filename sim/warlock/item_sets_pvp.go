package warlock

import (
	"time"

	"github.com/Vellasta/classic/sim/core"
	"github.com/Vellasta/classic/sim/core/stats"
)

var bluePvPBonuses = map[int32]core.ApplyEffect{
	// Increases damage and healing done by magical spells and effects by up to 23.
	2: func(agent core.Agent) {
		c := agent.GetCharacter()
		c.AddStat(stats.SpellPower, 23)
	},
	// Reduces the casting time of your Immolate spell by 0.2 sec.
	4: func(agent core.Agent) {
		warlock := agent.(WarlockAgent).GetWarlock()
		warlock.GetOrRegisterAura(core.Aura{
			Label: "Immolate Cast Time Reduction",
			OnInit: func(aura *core.Aura, sim *core.Simulation) {
				for _, spell := range warlock.Immolate {
					spell.DefaultCast.CastTime -= time.Millisecond * 200
					spell.DefaultCast.GCD -= time.Millisecond * 200
				}
			},
		})
	},
	// +20 Stamina.
	6: func(agent core.Agent) {
		c := agent.GetCharacter()
		c.AddStat(stats.Stamina, 20)
	},
}

var ItemSetChampionsThreads = core.NewItemSet(core.ItemSet{
	Name:    "Champion's Dreadgear",
	Bonuses: bluePvPBonuses,
})

var ItemSetLieutenantCommandersThreads = core.NewItemSet(core.ItemSet{
	Name:    "Lieutenant Commander's Dreadgear",
	Bonuses: bluePvPBonuses,
})

var epicPvpBonuses = map[int32]core.ApplyEffect{
	// +20 Stamina.
	2: func(agent core.Agent) {
		c := agent.GetCharacter()
		c.AddStat(stats.Stamina, 20)
	},
	// Reduces the casting time of your Immolate spell by 0.2 sec.
	3: func(agent core.Agent) {
		warlock := agent.(WarlockAgent).GetWarlock()
		warlock.GetOrRegisterAura(core.Aura{
			Label: "Immolate Cast Time Reduction",
			OnInit: func(aura *core.Aura, sim *core.Simulation) {
				for _, spell := range warlock.Immolate {
					spell.DefaultCast.CastTime -= time.Millisecond * 200
					spell.DefaultCast.GCD -= time.Millisecond * 200
				}
			},
		})
	},
	// Increases damage and healing done by magical spells and effects by up to 23.
	6: func(agent core.Agent) {
		c := agent.GetCharacter()
		c.AddStat(stats.SpellPower, 23)
	},
}

var ItemSetWarlordsThreads = core.NewItemSet(core.ItemSet{
	Name:    "Warlord's Threads",
	Bonuses: epicPvpBonuses,
})

var ItemSetFieldMarshalsThreads = core.NewItemSet(core.ItemSet{
	Name:    "Field Marshal's Threads",
	Bonuses: epicPvpBonuses,
})
