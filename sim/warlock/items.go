package warlock

import (
	"time"

	"github.com/Vellasta/classic/sim/core"
	"github.com/Vellasta/classic/sim/core/stats"
)

const (
	TheBlackBook                = 19337
	HazzarahsCharmOfDestruction = 19957
)

func init() {
	// https://www.wowhead.com/classic/item=19957/hazzarahs-charm-of-destruction
	// Use: Increases the critical hit chance of your Destruction spells by 10% for 20 sec. (3 Min Cooldown)
	core.NewItemEffect(HazzarahsCharmOfDestruction, func(agent core.Agent) {
		warlock := agent.(WarlockAgent).GetWarlock()

		actionID := core.ActionID{ItemID: HazzarahsCharmOfDestruction}
		duration := time.Second * 20

		var affectedSpells []*core.Spell

		buffAura := warlock.RegisterAura(core.Aura{
			ActionID: actionID,
			Label:    "Massive Destruction",
			Duration: duration,
			OnInit: func(aura *core.Aura, sim *core.Simulation) {
				affectedSpells = core.FilterSlice(warlock.Spellbook, func(spell *core.Spell) bool { return spell.Flags.Matches(WarlockFlagDestruction) })
			},
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				for _, spell := range affectedSpells {
					spell.BonusCritRating += 10 * core.SpellCritRatingPerCritChance
				}
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				for _, spell := range affectedSpells {
					spell.BonusCritRating -= 10 * core.SpellCritRatingPerCritChance
				}
			},
		})

		spell := warlock.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolFire,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    warlock.NewTimer(),
					Duration: time.Minute * 3,
				},
				SharedCD: core.Cooldown{
					Timer:    warlock.GetOffensiveTrinketCD(),
					Duration: duration,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		warlock.AddMajorCooldown(core.MajorCooldown{
			Spell:    spell,
			Priority: core.CooldownPriorityBloodlust,
			Type:     core.CooldownTypeDPS,
		})
	})

	// https://www.wowhead.com/classic/item=19337/the-black-book
	// Use: Empowers your pet, increasing pet damage by 100% and increasing pet armor by 100% for 30 sec.
	// This spell will only affect an Imp, Succubus, Incubus, Voidwalker, or Felhunter. (5 Min Cooldown)
	core.NewItemEffect(TheBlackBook, func(agent core.Agent) {
		warlock := agent.(WarlockAgent).GetWarlock()

		actionID := core.ActionID{ItemID: TheBlackBook}
		duration := time.Second * 30
		affectedPet := warlock.ActivePet

		statDeps := map[string]*stats.StatDependency{}
		for _, pet := range warlock.BasePets {
			statDeps[pet.Name] = pet.NewDynamicMultiplyStat(stats.Armor, 2)
		}

		buffAura := warlock.RegisterAura(core.Aura{
			ActionID: actionID,
			Label:    "Blessing of the Black Book",
			Duration: duration,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				affectedPet = warlock.ActivePet
				if affectedPet != nil {
					affectedPet.PseudoStats.DamageDealtMultiplier *= 2.0
					affectedPet.EnableDynamicStatDep(sim, statDeps[affectedPet.Name])
				}
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				if affectedPet != nil {
					affectedPet.PseudoStats.DamageDealtMultiplier /= 2.0
					affectedPet.DisableDynamicStatDep(sim, statDeps[affectedPet.Name])
				}
			},
		})

		spell := warlock.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolShadow,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    warlock.NewTimer(),
					Duration: time.Minute * 5,
				},
				SharedCD: core.Cooldown{
					Timer:    warlock.GetOffensiveTrinketCD(),
					Duration: duration,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		warlock.AddMajorCooldown(core.MajorCooldown{
			Spell:    spell,
			Priority: core.CooldownPriorityBloodlust,
			Type:     core.CooldownTypeDPS,
		})
	})
}
