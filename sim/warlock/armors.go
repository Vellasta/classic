package warlock

import (
	"github.com/Vellasta/classic/sim/core"
	"github.com/Vellasta/classic/sim/core/stats"
)

func (warlock *Warlock) applyDemonArmor() {
	spellID := map[int32]int32{
		25: 706,
		40: 11733,
		50: 11734,
		60: 11735,
	}[warlock.Level]

	armor := map[int32]float64{
		25: 210.0,
		40: 390.0,
		50: 480.0,
		60: 570.0,
	}[warlock.Level]

	shadowRes := map[int32]float64{
		25: 3.0,
		40: 9.0,
		50: 12.0,
		60: 15.0,
	}[warlock.Level]

	warlock.AddStat(stats.Armor, armor)
	warlock.AddStat(stats.ShadowResistance, shadowRes)

	warlock.GetOrRegisterAura(core.Aura{
		Label:    "Demon Armor",
		ActionID: core.ActionID{SpellID: spellID},
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
	})
}
