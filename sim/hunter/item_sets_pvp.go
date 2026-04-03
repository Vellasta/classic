package hunter

import (
	"time"

	"github.com/Vellasta/classic/sim/core"
	"github.com/Vellasta/classic/sim/core/stats"
)

///////////////////////////////////////////////////////////////////////////
//                            Classic Phase 2
///////////////////////////////////////////////////////////////////////////

// https://www.wowhead.com/classic/item-set=362/lieutenant-commanders-pursuit
var ItemSetLieutenantCommandersPursuit = core.NewItemSet(core.ItemSet{
	Name: "Lieutenant Commander's Pursuit",
	Bonuses: map[int32]core.ApplyEffect{
		// Increases your chance to parry an attack by 1%.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Parry, 1*core.ParryRatingPerParryChance)
		},
		// Reduces the cooldown of your Concussive Shot by 1 sec.
		4: func(agent core.Agent) {
			// Nothing to do
		},
		// +15 Stamina.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 15)
		},
	},
})

// https://www.wowhead.com/classic/item-set=361/champions-pursuit
var ItemSetChampionsPursuit = core.NewItemSet(core.ItemSet{
	Name: "Champion's Pursuit",
	Bonuses: map[int32]core.ApplyEffect{
		// Increases your chance to parry an attack by 1%.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Parry, 1*core.ParryRatingPerParryChance)
		},
		// Reduces the cooldown of your Concussive Shot by 1 sec.
		4: func(agent core.Agent) {
			// Nothing to do
		},
		// +15 Stamina.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 15)
		},
	},
})

// https://www.wowhead.com/classic/item-set=543/champions-pursuance
var ItemSetChampionsPursuance = core.NewItemSet(core.ItemSet{
	Name: "Champion's Pursuance",
	Bonuses: map[int32]core.ApplyEffect{
		// +20 Agility.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Agility, 20)
		},
		// Reduces the cooldown of your Concussive Shot by 1 sec.
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

// https://www.wowhead.com/classic/item-set=550/lieutenant-commanders-pursuance
var ItemSetLieutenantCommandersPursuance = core.NewItemSet(core.ItemSet{
	Name: "Lieutenant Commander's Pursuance",
	Bonuses: map[int32]core.ApplyEffect{
		// +20 Agility.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Agility, 20)
		},
		// Reduces the cooldown of your Concussive Shot by 1 sec.
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
//                            Classic Phase 3
///////////////////////////////////////////////////////////////////////////

var ItemSetWarlordsPursuit = core.NewItemSet(core.ItemSet{
	Name: "Warlord's Pursuit",
	Bonuses: map[int32]core.ApplyEffect{
		// 20 Stamina
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
		// Reduces the cooldown of your Concussive Shot by 1 sec.
		4: func(agent core.Agent) {
			// Nothing to do
		},
		// +20 Agi
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Agility, 20)
		},
	},
})

var ItemSetFieldMarshalsPursuit = core.NewItemSet(core.ItemSet{
	Name: "Field Marshal's Pursuit",
	Bonuses: map[int32]core.ApplyEffect{
		// 20 stamina
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
		// Reduces the cooldown of your Concussive Shot by 1 sec.
		4: func(agent core.Agent) {
			// Nothing to do
		},
		// +20 Agi
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Agility, 20)
		},
	},
})

///////////////////////////////////////////////////////////////////////////
//                            New PvP Sets
///////////////////////////////////////////////////////////////////////////

var ItemSetCombatantsPursuit = core.NewItemSet(core.ItemSet{
	Name: "Combatant's Pursuit",
	Bonuses: map[int32]core.ApplyEffect{
		// (2) Set: Increases the damage done by your Multi-Shot and Carve by 4%.
		2: func(agent core.Agent) {
			hunter := agent.(HunterAgent).GetHunter()
			hunter.RegisterAura(core.Aura{
				Label: "Combatant's Multishot And Carve Damage Increase",
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					hunter.MultiShot.BaseDamageMultiplierAdditive += 0.04
					if hunter.Carve != nil {
						hunter.Carve.BaseDamageMultiplierAdditive += 0.04
					}
				},
			})
		},
		// (4) Set: Reduces the cooldown of Concussive Shot by 1000 sec and Wing Clip by 500 sec.
		4: func(agent core.Agent) {
			hunter := agent.(HunterAgent).GetHunter()
			core.MakePermanent(hunter.RegisterAura(core.Aura{
				Label: "Combatant's Wing Clip Cooldown",
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					hunter.WingClip.CD.Duration -= time.Millisecond * 500
				},
			}))
		},
		// (6) Set: +20 Stamina.
		// (6) Set: Reduces damage taken from critical hits and damage over time effects by 3%.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
	},
})

var ItemSetPartisansPursuit = core.NewItemSet(core.ItemSet{
	Name: "Partisan's Pursuit",
	Bonuses: map[int32]core.ApplyEffect{
		// (2) Set: Increases the damage done by your Multi-Shot and Carve by 4%.
		2: func(agent core.Agent) {
			hunter := agent.(HunterAgent).GetHunter()
			hunter.RegisterAura(core.Aura{
				Label: "Partisan's Multishot And Carve Damage Increase",
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					hunter.MultiShot.BaseDamageMultiplierAdditive += 0.04
					if hunter.Carve != nil {
						hunter.Carve.BaseDamageMultiplierAdditive += 0.04
					}
				},
			})
		},
		// (4) Set: Reduces the cooldown of Concussive Shot by 1000 sec and Wing Clip by 500 sec.
		4: func(agent core.Agent) {
			hunter := agent.(HunterAgent).GetHunter()
			core.MakePermanent(hunter.RegisterAura(core.Aura{
				Label: "Partisan's Wing Clip Cooldown",
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					hunter.WingClip.CD.Duration -= time.Millisecond * 500
				},
			}))
		},
		// (6) Set: +30 Stamina.
		// (6) Set: Reduces damage taken from critical hits and damage over time effects by 3%.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 30)
		},
	},
})

var ItemSetVeteransPursuit = core.NewItemSet(core.ItemSet{
	Name: "Veteran's Pursuit",
	Bonuses: map[int32]core.ApplyEffect{
		// (2) Set: Increases the damage done by your Multi-Shot and Carve by 4%.
		2: func(agent core.Agent) {
			hunter := agent.(HunterAgent).GetHunter()
			hunter.RegisterAura(core.Aura{
				Label: "Veteran's Multishot And Carve Damage Increase",
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					hunter.MultiShot.BaseDamageMultiplierAdditive += 0.04
					if hunter.Carve != nil {
						hunter.Carve.BaseDamageMultiplierAdditive += 0.04
					}
				},
			})
		},
		// (4) Set: Reduces the cooldown of Concussive Shot by 1000 sec and Wing Clip by 500 sec.
		4: func(agent core.Agent) {
			hunter := agent.(HunterAgent).GetHunter()
			core.MakePermanent(hunter.RegisterAura(core.Aura{
				Label: "Veteran's Wing Clip Cooldown",
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					hunter.WingClip.CD.Duration -= time.Millisecond * 500
				},
			}))
		},
		// (6) Set: +35 Stamina.
		// (6) Set: Reduces damage taken from critical hits and damage over time effects by 6%.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 35)
		},
	},
})
