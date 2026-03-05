package rogue

import (
	"github.com/Vellasta/classic/sim/core"
	"github.com/Vellasta/classic/sim/core/stats"
)

///////////////////////////////////////////////////////////////////////////
//                          Classic Phase 2
///////////////////////////////////////////////////////////////////////////

var ItemSetChampionsGuard = core.NewItemSet(core.ItemSet{
	Name: "Champion's Guard",
	Bonuses: map[int32]core.ApplyEffect{
		// +40 Attack Power.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStats(stats.Stats{
				stats.AttackPower:       40,
				stats.RangedAttackPower: 40,
			})
		},
		// Reduces the cooldown of your Gouge ability by 1 sec.
		4: func(agent core.Agent) {
			// Nothing to do
		},
		// +20 Stamina.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
	},
})

var ItemSetLieutenantCommandersGuard = core.NewItemSet(core.ItemSet{
	Name: "Lieutenant Commander's Guard",
	Bonuses: map[int32]core.ApplyEffect{
		// +40 Attack Power.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStats(stats.Stats{
				stats.AttackPower:       40,
				stats.RangedAttackPower: 40,
			})
		},
		// Reduces the cooldown of your Gouge ability by 1 sec.
		4: func(agent core.Agent) {
			// Nothing to do
		},
		// +20 Stamina.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
	},
})

var ItemSetWarlordsVestments = core.NewItemSet(core.ItemSet{
	Name: "Warlord's Vestments",
	Bonuses: map[int32]core.ApplyEffect{
		// +20 Stamina.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
		// Reduces the cooldown of your Gouge ability by 1 sec.
		4: func(agent core.Agent) {
			// Nothing to do
		},
		// +40 Attack Power.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStats(stats.Stats{
				stats.AttackPower:       40,
				stats.RangedAttackPower: 40,
			})
		},
	},
})

var ItemSetFieldMarshalsVestments = core.NewItemSet(core.ItemSet{
	Name: "Field Marshal's Vestments",
	Bonuses: map[int32]core.ApplyEffect{
		// +20 Stamina.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
		// Reduces the cooldown of your Gouge ability by 1 sec.
		4: func(agent core.Agent) {
			// Nothing to do
		},
		// +40 Attack Power.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStats(stats.Stats{
				stats.AttackPower:       40,
				stats.RangedAttackPower: 40,
			})
		},
	},
})
