package mage

import (
	"github.com/Vellasta/classic/sim/core"
	"github.com/Vellasta/classic/sim/core/stats"
)

///////////////////////////////////////////////////////////////////////////
//                            Clasic Phase 2
///////////////////////////////////////////////////////////////////////////

var ItemSetChampionsArcanum = core.NewItemSet(core.ItemSet{
	Name: "Champion's Arcanum",
	Bonuses: map[int32]core.ApplyEffect{
		// Increases damage and healing done by magical spells and effects by up to 23.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.SpellPower, 23)
		},
		// Reduces the cooldown of your Blink spell by 1.5 sec.
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

var ItemSetLieutenantCommandersArcanum = core.NewItemSet(core.ItemSet{
	Name: "Lieutenant Commander's Arcanum",
	Bonuses: map[int32]core.ApplyEffect{
		// Increases damage and healing done by magical spells and effects by up to 23.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.SpellPower, 23)
		},
		// Reduces the cooldown of your Blink spell by 1.5 sec.
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

///////////////////////////////////////////////////////////////////////////
//                            Phase 3 Items
///////////////////////////////////////////////////////////////////////////

var ItemSetWarlordsRegalia = core.NewItemSet(core.ItemSet{
	Name: "Warlord's Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		// +20 Stamina.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
		// Reduces the cooldown of your Blink spell by 1.5 sec.
		4: func(agent core.Agent) {
			// Nothing to do
		},
		// Increases damage and healing done by magical spells and effects by up to 23.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.SpellPower, 23)
		},
	},
})

var ItemSetFieldMarshalsRegalia = core.NewItemSet(core.ItemSet{
	Name: "Field Marshal's Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		// +20 Stamina.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
		// Reduces the cooldown of your Blink spell by 1.5 sec.
		4: func(agent core.Agent) {
			// Nothing to do
		},
		// Increases damage and healing done by magical spells and effects by up to 23.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.SpellPower, 23)
		},
	},
})
