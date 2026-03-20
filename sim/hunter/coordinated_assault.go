package hunter

import (
	"time"

	"github.com/Vellasta/classic/sim/core"
)

func (hunter *Hunter) newCoordinatedAssault() *core.Spell {
	return hunter.RegisterSpell(core.SpellConfig{
		SpellCode:   SpellCode_HunterCoordinatedAssault,
		ActionID:    core.ActionID{SpellID: 49557},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagMeleeMetrics,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			damage := 0.2 * spell.MeleeAttackPower(target)
			spell.CalcAndDealDamage(sim, target, damage, spell.OutcomeMeleeSpecialHitAndCrit)
		},
	})
}

func (hp *HunterPet) createCoordinatedAssaultAura(auraLabel string, actionID core.ActionID) *core.Aura {
	return hp.GetOrRegisterAura(core.Aura{
		Label:    auraLabel,
		ActionID: actionID,
		Duration: time.Second * 6,

		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.ProcMask.Matches(core.ProcMaskMeleeWhiteHit) {
				hp.hunterOwner.CoordinatedAssault.Cast(sim, result.Target)
				aura.Deactivate(sim)
			}
		},
	})
}

func (hunter *Hunter) applyCoordinatedAssault() {
	if !hunter.Talents.CoordinatedAssault {
		return
	}

	icd := core.Cooldown{
		Timer:    hunter.NewTimer(),
		Duration: time.Second * 3,
	}

	hunter.CoordinatedAssault = hunter.newCoordinatedAssault()

	hunter.pet.coordinatedAssaultAura = hunter.pet.createCoordinatedAssaultAura(
		"Strike Together",
		core.ActionID{SpellID: 49558},
	)

	core.MakePermanent(hunter.RegisterAura(core.Aura{
		Label: "Coordinated Assault Talent",
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !(spell.SpellCode == SpellCode_HunterArcaneShot || spell.SpellCode == SpellCode_HunterSteadyShot || spell.SpellCode == SpellCode_HunterRaptorStrikeHit) {
				return
			}

			if hunter.pet.coordinatedAssaultAura != nil && icd.IsReady(sim) {
				icd.Use(sim)
				hunter.pet.coordinatedAssaultAura.Activate(sim)
			}
		},
	}))
}
