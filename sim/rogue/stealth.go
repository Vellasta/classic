package rogue

import (
	"github.com/Vellasta/classic/sim/core"
)

func (rogue *Rogue) registerStealthAura() {
	rogue.StealthAura = rogue.RegisterAura(core.Aura{
		Label:    "Stealth",
		ActionID: core.ActionID{SpellID: 1787},
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			// Stealth triggered auras
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
		},
		// Stealth breaks on damage taken (if not absorbed)
		// This may be desirable later, but not applicable currently
	})
}
