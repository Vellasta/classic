package warrior

import (
	"time"

	"github.com/Vellasta/classic/sim/core"
)

const ShoutExpirationThreshold = time.Second * 3

func (warrior *Warrior) newShoutSpellConfig(actionID core.ActionID, rank int32, allyAuras core.AuraArray) *WarriorSpell {
	// Use extra hits to simulate buffing your party for threat
	extraHits := 5 - len(allyAuras)

	return warrior.RegisterSpell(AnyStance, core.SpellConfig{
		ActionID: actionID,
		Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagAPL | core.SpellFlagHelpful,

		RageCost: core.RageCostOptions{
			Cost: 10,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
		},

		FlatThreatBonus: float64(core.BattleShoutLevel[rank]),

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			for _, aura := range allyAuras {
				spell.CalcAndDealOutcome(sim, aura.Unit, spell.OutcomeAlwaysHit)
				aura.Activate(sim)
			}

			for i := 0; i < extraHits; i++ {
				spell.CalcAndDealOutcome(sim, target, spell.OutcomeAlwaysHit)
			}
		},

		RelatedAuras: []core.AuraArray{allyAuras},
	})
}

func (warrior *Warrior) registerBattleShout() {
	rank := core.TernaryInt32(core.IncludeAQ, 7, 6)
	actionId := core.BattleShoutSpellId[rank]
	has3pcWrath := warrior.HasSetBonus(ItemSetBattleGearOfWrath, 3)

	warrior.BattleShout = warrior.newShoutSpellConfig(core.ActionID{SpellID: actionId}, rank, warrior.NewPartyAuraArray(func(unit *core.Unit) *core.Aura {
		return core.BattleShoutAura(unit, warrior.Talents.ImprovedBattleShout, warrior.Talents.BoomingVoice, has3pcWrath)
	}))
}

func (warrior *Warrior) registerShouts() {
	warrior.registerBattleShout()
}
