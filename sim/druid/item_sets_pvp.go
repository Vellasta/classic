package druid

import (
	"github.com/Vellasta/classic/sim/core"
	"github.com/Vellasta/classic/sim/core/stats"
)

var ItemSetChampionsRefuge = core.NewItemSet(core.ItemSet{
	Name: "Champion's Refuge",
	Bonuses: map[int32]core.ApplyEffect{
		// +40 Attack Power.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.AttackPower, 40)
		},
		// Increases your movement speed by 15% while in Bear, Cat, or Travel Form. Only active outdoors.
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

var ItemSetLieutenantCommandersRefuge = core.NewItemSet(core.ItemSet{
	Name: "Lieutenant Commander's Refuge",
	Bonuses: map[int32]core.ApplyEffect{
		// +40 Attack Power.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.AttackPower, 40)
		},
		// Increases your movement speed by 15% while in Bear, Cat, or Travel Form. Only active outdoors.
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

var ItemSetFieldMarshalsSanctuary = core.NewItemSet(core.ItemSet{
	Name: "Field Marshal's Sanctuary",
	Bonuses: map[int32]core.ApplyEffect{
		// +20 Stamina.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
		// Increases your movement speed by 15% while in Bear, Cat, or Travel Form. Only active outdoors.
		3: func(agent core.Agent) {
			// Nothing to do
		},
		// +40 Attack Power.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.AttackPower, 40)
			c.AddStat(stats.RangedAttackPower, 40)
		},
	},
})

var ItemSetWarlordsSanctuary = core.NewItemSet(core.ItemSet{
	Name: "Warlord's Sanctuary",
	Bonuses: map[int32]core.ApplyEffect{
		// +20 Stamina.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
		// Increases your movement speed by 15% while in Bear, Cat, or Travel Form. Only active outdoors.
		3: func(agent core.Agent) {
			// Nothing to do
		},
		// +40 Attack Power.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.AttackPower, 40)
			c.AddStat(stats.RangedAttackPower, 40)
		},
	},
})
