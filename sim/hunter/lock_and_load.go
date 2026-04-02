package hunter

import (
	"time"

	"github.com/Vellasta/classic/sim/core"
)

// Utility function to create an Improved Hawk Aura
func (hunter *Hunter) createLockAndLoadAura(auraLabel string, actionID core.ActionID) *core.Aura {
	return hunter.GetOrRegisterAura(core.Aura{
		Label:    auraLabel,
		ActionID: actionID,
		Duration: time.Second * 10,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			hunter.AimedShot.DefaultCast.CastTime = time.Duration(int(float64(hunter.AimedShot.DefaultCast.CastTime.Nanoseconds()) * float64(1/2.0)))
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			hunter.AimedShot.DefaultCast.CastTime = time.Duration(int(float64(hunter.AimedShot.DefaultCast.CastTime.Nanoseconds()) * float64(2.0)))
		},
	})
}

func (hunter *Hunter) applyLockAndLoad() {
	if !hunter.Talents.LockAndLoad {
		return
	}

	hunter.lockAndLoadAura = hunter.createLockAndLoadAura(
		"Lock and Load",
		core.ActionID{SpellID: 52920},
	)

	core.MakePermanent(hunter.RegisterAura(core.Aura{
		Label: "Lock and Load Talent",
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.ProcMask.Matches(core.ProcMaskEmpty) || !(spell.SpellCode == SpellCode_HunterSteadyShot || spell.SpellCode == SpellCode_HunterAimedShot || spell.SpellCode == SpellCode_HunterArcaneShot) {
				return
			}

			if result.DidCrit() {
				hunter.AimedShot.CD.Reset()
				if !hunter.lockAndLoadAura.IsActive() {
					hunter.lockAndLoadAura.Activate(sim)
				}
			}
		},
	}))
}
