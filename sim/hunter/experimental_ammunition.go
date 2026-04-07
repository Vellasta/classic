package hunter

import (
	"time"

	"github.com/Vellasta/classic/sim/core"
	"github.com/Vellasta/classic/sim/core/proto"
)

func (hunter *Hunter) createExplosiveAmmunitionAura(auraLabel string, actionID core.ActionID) *core.Aura {
	return hunter.GetOrRegisterAura(core.Aura{
		Label:    auraLabel,
		ActionID: actionID,
		Duration: time.Second * 60,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			hunter.PoisonousAmmunitionAura.Deactivate(sim)
			hunter.EnchantedAmmunitionAura.Deactivate(sim)
		},
	})
}

func (hunter *Hunter) createPoisonousAmmunitionAura(auraLabel string, actionID core.ActionID) *core.Aura {
	return hunter.GetOrRegisterAura(core.Aura{
		Label:    auraLabel,
		ActionID: actionID,
		Duration: time.Second * 60,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			hunter.ExplosiveAmmunitionAura.Deactivate(sim)
			hunter.EnchantedAmmunitionAura.Deactivate(sim)
		},
	})
}

func (hunter *Hunter) createEnchantedAmmunitionAura(auraLabel string, actionID core.ActionID) *core.Aura {
	return hunter.GetOrRegisterAura(core.Aura{
		Label:    auraLabel,
		ActionID: actionID,
		Duration: time.Second * 60,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			hunter.ExplosiveAmmunitionAura.Deactivate(sim)
			hunter.PoisonousAmmunitionAura.Deactivate(sim)
		},
	})
}

func (hunter *Hunter) applyExplosiveAmmunition() {
	hunter.ExplosiveAmmunition = hunter.RegisterSpell(core.SpellConfig{
		SpellCode:    SpellCode_HunterExplosiveAmmunition,
		ActionID:     core.ActionID{SpellID: 58109},
		SpellSchool:  core.SpellSchoolFire,
		DefenseType:  core.DefenseTypeRanged,
		ProcMask:     core.ProcMaskSpellDamage,
		Flags:        core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell | SpellFlagShot,
		CastType:     proto.CastType_CastTypeRanged,
		MissileSpeed: 24,

		CritDamageBonus: hunter.mortalShots(),

		DamageMultiplier: 0.05 * (1 + 0.05*float64(hunter.Talents.ImprovedMarksmanship)),
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := hunter.AutoAttacks.Ranged().CalculateNormalizedWeaponDamage(sim, spell.RangedAttackPower(target, false)) +
				hunter.NormalizedAmmoDamageBonus +
				600

			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeRangedHitAndCrit)
			spell.DealDamage(sim, result)
			hunter.ExplosiveAmmunitionAura.Activate(sim)
			hunter.nextExperimentalAmmunitionState()
		},
	})
}

func (hunter *Hunter) applyPoisonousAmmunition() {
	hunter.PoisonousAmmunition = hunter.RegisterSpell(core.SpellConfig{
		SpellCode:    SpellCode_HunterPoisonousAmmunition,
		ActionID:     core.ActionID{SpellID: 58110},
		SpellSchool:  core.SpellSchoolNature,
		DefenseType:  core.DefenseTypeRanged,
		ProcMask:     core.ProcMaskSpellDamage,
		Flags:        core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell | SpellFlagShot,
		CastType:     proto.CastType_CastTypeRanged,
		MissileSpeed: 24,

		CritDamageBonus: hunter.mortalShots(),

		DamageMultiplier: 0.05 * (1 + 0.05*float64(hunter.Talents.ImprovedMarksmanship)),
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := hunter.AutoAttacks.Ranged().CalculateNormalizedWeaponDamage(sim, spell.RangedAttackPower(target, false)) +
				hunter.NormalizedAmmoDamageBonus +
				600

			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeRangedHitAndCrit)
			spell.DealDamage(sim, result)
			hunter.PoisonousAmmunitionAura.Activate(sim)
			hunter.nextExperimentalAmmunitionState()
		},
	})
}

func (hunter *Hunter) applyEnchantedAmmunition() {
	hunter.EnchantedAmmunition = hunter.RegisterSpell(core.SpellConfig{
		SpellCode:    SpellCode_HunterEnchantedAmmunition,
		ActionID:     core.ActionID{SpellID: 58111},
		SpellSchool:  core.SpellSchoolArcane,
		DefenseType:  core.DefenseTypeRanged,
		ProcMask:     core.ProcMaskSpellDamage,
		Flags:        core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell | SpellFlagShot,
		CastType:     proto.CastType_CastTypeRanged,
		MissileSpeed: 24,

		CritDamageBonus: hunter.mortalShots(),

		DamageMultiplier: 0.05 * (1 + 0.05*float64(hunter.Talents.ImprovedMarksmanship)),
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := hunter.AutoAttacks.Ranged().CalculateNormalizedWeaponDamage(sim, spell.RangedAttackPower(target, false)) +
				hunter.NormalizedAmmoDamageBonus +
				600

			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeRangedHitAndCrit)
			spell.DealDamage(sim, result)
			hunter.EnchantedAmmunitionAura.Activate(sim)
			hunter.nextExperimentalAmmunitionState()
		},
	})
}

func (hunter *Hunter) applyExplosiveAmmunitionAOE() {
	hunter.ExplosiveAmmunitionAOE = hunter.RegisterSpell(core.SpellConfig{
		SpellCode:   SpellCode_HunterExplosiveAmmunitionAOE,
		ActionID:    core.ActionID{SpellID: 58112},
		SpellSchool: core.SpellSchoolFire,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell | SpellFlagShot,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			curTarget := target
			baseDamage := 0.2 * (hunter.AutoAttacks.Ranged().CalculateNormalizedWeaponDamage(sim, spell.RangedAttackPower(target, false)) +
				hunter.NormalizedAmmoDamageBonus)
			baseDamage *= sim.Encounter.AOECapMultiplier()

			numHits := hunter.Env.GetNumTargets()
			for hitIndex := int32(0); hitIndex < numHits; hitIndex++ {
				spell.CalcAndDealDamage(sim, curTarget, baseDamage, spell.OutcomeRangedHit)
				curTarget = sim.Environment.NextTargetUnit(curTarget)
			}
		},
	})
}

func (hunter *Hunter) registerExperimentalAmmunition() {
	if !hunter.Talents.ExperimentalAmmunition {
		return
	}

	hunter.applyExplosiveAmmunition()
	hunter.applyPoisonousAmmunition()
	hunter.applyEnchantedAmmunition()

	hunter.ExplosiveAmmunitionAura = hunter.createExplosiveAmmunitionAura(
		"Explosive Ammunition Aura",
		core.ActionID{SpellID: 58109},
	)
	hunter.PoisonousAmmunitionAura = hunter.createPoisonousAmmunitionAura(
		"Poisonous Ammunition Aura",
		core.ActionID{SpellID: 58110},
	)
	hunter.EnchantedAmmunitionAura = hunter.createEnchantedAmmunitionAura(
		"Enchanted Ammunition Aura",
		core.ActionID{SpellID: 58111},
	)

	hunter.applyExplosiveAmmunitionAOE()
	hunter.PoisonousAmmunitionDebuff = hunter.NewEnemyAuraArray(core.PoisonousAmmunitionAura)
	hunter.EnchantedAmmunitionDebuff = hunter.NewEnemyAuraArray(core.EnchantedAmmunitionAura)

	hunter.ExperimentalAmmunitionTypes = append(hunter.ExperimentalAmmunitionTypes, hunter.ExplosiveAmmunition)
	hunter.ExperimentalAmmunitionTypes = append(hunter.ExperimentalAmmunitionTypes, hunter.PoisonousAmmunition)
	hunter.ExperimentalAmmunitionTypes = append(hunter.ExperimentalAmmunitionTypes, hunter.EnchantedAmmunition)

	hunter.ExperimentalAmmunitionState = 0
}

func (hunter *Hunter) nextExperimentalAmmunitionState() {
	hunter.ExperimentalAmmunitionState = (hunter.ExperimentalAmmunitionState + 1) % 3
}
