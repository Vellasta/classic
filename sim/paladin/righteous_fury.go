package paladin

import (
	"github.com/Vellasta/classic/sim/core"
)

func (paladin *Paladin) registerRighteousFury() {
	if !paladin.Options.RighteousFury {
		return
	}
	actionID := core.ActionID{SpellID: 25780}

	// Improved Righteous Fury is multiplicative.
	rfThreatMultiplier := 1.6 * (1 + []float64{0.0, 0.16, 0.33, 0.5}[paladin.Talents.ImprovedRighteousFury])

	paladin.OnSpellRegistered(func(spell *core.Spell) {
		if spell.SpellSchool.Matches(core.SpellSchoolHoly) {
			spell.ThreatMultiplier *= rfThreatMultiplier
		}
	})

	rfAura := core.MakePermanent(&core.Aura{Label: "Righteous Fury", ActionID: actionID})
	paladin.RegisterAura(*rfAura)
}
