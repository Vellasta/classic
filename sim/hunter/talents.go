package hunter

import (
	"time"

	"github.com/Vellasta/classic/sim/core"
	"github.com/Vellasta/classic/sim/core/proto"
	"github.com/Vellasta/classic/sim/core/stats"
)

func (hunter *Hunter) ApplyTalents() {
	if hunter.pet != nil {
		hunter.applyFrenzy()
		hunter.applyCoordinatedAssault()
		hunter.registerBestialWrathCD()
		hunter.registerScentOfBlood()

		hunter.pet.AddStat(stats.MeleeCrit, core.CritRatingPerCritChance*3*float64(hunter.Talents.Ferocity))
		hunter.pet.AddStat(stats.SpellCrit, core.SpellCritRatingPerCritChance*3*float64(hunter.Talents.Ferocity))

		hunter.pet.AddStat(stats.MeleeHit, 4*float64(hunter.Talents.BestialPrecision))
		hunter.pet.AddStat(stats.SpellHit, 9*float64(hunter.Talents.BestialPrecision))
		hunter.pet.PseudoStats.FeralCombatEnabled = true
		hunter.pet.PseudoStats.FeralCombatSkill += 5 * float64(hunter.Talents.BestialPrecision)

		hunter.pet.PseudoStats.DamageDealtMultiplier *= 1 + 0.04*float64(hunter.Talents.UnleashedFury)

		if hunter.Talents.EnduranceTraining > 0 {
			hunter.pet.MultiplyStat(stats.Health, 1+(0.03*float64(hunter.Talents.EnduranceTraining)))
		}
	}

	if hunter.Talents.ImprovedSlaying > 0 {
		hunter.Env.RegisterPostFinalizeEffect(func() {
			for _, t := range hunter.Env.Encounter.Targets {
				switch t.MobType {
				case proto.MobType_MobTypeHumanoid:
					multiplier := []float64{1, 1.01, 1.02, 1.03}[hunter.Talents.ImprovedSlaying]
					for _, at := range hunter.AttackTables[t.UnitIndex] {
						at.DamageDealtMultiplier *= multiplier
						at.CritMultiplier *= multiplier
					}
				case proto.MobType_MobTypeBeast, proto.MobType_MobTypeGiant, proto.MobType_MobTypeDragonkin:
					multiplier := []float64{1, 1.01, 1.02, 1.03}[hunter.Talents.ImprovedSlaying]
					for _, at := range hunter.AttackTables[t.UnitIndex] {
						at.DamageDealtMultiplier *= multiplier
						at.CritMultiplier *= multiplier
					}
				}
			}
		})
	}

	if hunter.Talents.AloneAgainstTheWorld > 0 {
		hunter.PseudoStats.DamageDealtMultiplier *= 1 + 0.03*float64(hunter.Talents.AloneAgainstTheWorld)
	}

	if hunter.Talents.BestialDiscipline > 0 {
		core.MakePermanent(hunter.RegisterAura(core.Aura{
			Label: "Bestial Discipline",
			OnInit: func(aura *core.Aura, sim *core.Simulation) {
				if hunter.pet != nil {
					hunter.pet.AddFocusRegenMultiplier(0.1 * float64(hunter.Talents.BestialDiscipline))
					reducedCDFraction := float64(1.0 - (0.1 * float64(hunter.Talents.BestialDiscipline) / 1.0))
					if hunter.pet.specialAbility != nil {
						hunter.pet.specialAbility.CD.Duration = time.Duration(int(float64(hunter.pet.specialAbility.CD.Duration.Nanoseconds()) * reducedCDFraction))
					}
					if hunter.pet.focusDump != nil {
						hunter.pet.focusDump.CD.Duration = time.Duration(int(float64(hunter.pet.focusDump.CD.Duration.Nanoseconds()) * reducedCDFraction))
					}
				}
			},
		}))
	}

	hunter.AddStat(stats.MeleeHit, float64(hunter.Talents.Surefooted)*1*core.MeleeHitRatingPerHitChance)
	if hunter.AutoAttacks.IsDualWielding {
		hunter.AddStat(stats.MeleeHit, float64(hunter.Talents.Surefooted)*1*core.MeleeHitRatingPerHitChance)
	}
	hunter.AddStat(stats.SpellHit, float64(hunter.Talents.Surefooted)*1*core.SpellHitRatingPerHitChance)

	hunter.AddStat(stats.MeleeCrit, float64(hunter.Talents.KillerInstinct)*1*core.CritRatingPerCritChance)

	if hunter.Talents.LethalShots > 0 {
		lethalBonus := 1 * float64(hunter.Talents.LethalShots) * core.CritRatingPerCritChance
		hunter.OnSpellRegistered(func(spell *core.Spell) {
			if spell.Flags.Matches(SpellFlagShot) {
				spell.BonusCritRating += lethalBonus
			}
		})
		hunter.AutoAttacks.RangedConfig().BonusCritRating += lethalBonus
	}

	if hunter.Talents.RangedWeaponSpecialization > 0 {
		mult := 1 + 0.02*float64(hunter.Talents.RangedWeaponSpecialization)
		hunter.OnSpellRegistered(func(spell *core.Spell) {
			if spell.ProcMask.Matches(core.ProcMaskRanged) && spell.SpellCode != SpellCode_HunterSerpentSting {
				spell.DamageMultiplier *= mult
			}
		})
	}

	if hunter.Talents.SwiftReflexes > 0 {
		hunter.PseudoStats.MeleeSpeedMultiplier *= 1.0 + 0.01*float64(hunter.Talents.SwiftReflexes)
		hunter.PseudoStats.RangedSpeedMultiplier *= 1.0 + 0.01*float64(hunter.Talents.SwiftReflexes)
	}

	if hunter.Talents.SavageStrikes > 0 {
		hunter.AutoAttacks.OHConfig().DamageMultiplier *= 1.0 + 0.125*float64(hunter.Talents.SavageStrikes)
	}

	if hunter.Talents.Survivalist > 0 {
		hunter.MultiplyStat(stats.Health, 1.0+0.02*float64(hunter.Talents.Survivalist))
	}

	if hunter.Talents.KillerInstinct > 0 {
		hunter.AutoAttacks.MHConfig().CritDamageBonus = 0.066 * float64(hunter.Talents.KillerInstinct)
		hunter.AutoAttacks.OHConfig().CritDamageBonus = 0.066 * float64(hunter.Talents.KillerInstinct)
	}

	if hunter.Talents.LightningReflexes > 0 {
		agiBonus := 0.02 * float64(hunter.Talents.LightningReflexes)
		hunter.MultiplyStat(stats.Agility, 1.0+agiBonus)
	}

	hunter.applyEfficiency()
	hunter.applyTrapMastery()
	hunter.applyPiercingShots()
	hunter.applyEndlessQuiver()
	hunter.applyLockAndLoad()
}

func (hunter *Hunter) applyFrenzy() {
	if hunter.Talents.Frenzy == 0 {
		return
	}

	procChance := 0.2 * float64(hunter.Talents.Frenzy)

	procAura := hunter.pet.RegisterAura(core.Aura{
		Label:    "Frenzy Proc",
		ActionID: core.ActionID{SpellID: 19625},
		Duration: time.Second * 8,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.MultiplyAttackSpeed(sim, 1.3)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.MultiplyAttackSpeed(sim, 1/1.3)
		},
	})

	hunter.pet.RegisterAura(core.Aura{
		Label:    "Frenzy",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellResult *core.SpellResult) {
			if !spellResult.Outcome.Matches(core.OutcomeCrit) {
				return
			}
			if procChance == 1 || sim.RandomFloat("Frenzy") < procChance {
				procAura.Activate(sim)
			}
		},
	})
}

func (hunter *Hunter) registerBestialWrathCD() {
	if !hunter.Talents.BestialWrath {
		return
	}

	actionID := core.ActionID{SpellID: 19574}

	bwSpell := hunter.RegisterSpell(core.SpellConfig{
		ActionID: actionID,
		Flags:    core.SpellFlagAPL,

		ManaCost: core.ManaCostOptions{
			BaseCost: 0.12,
		},

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    hunter.NewTimer(),
				Duration: time.Second * 90,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			hunter.pet.ScentOfBloodPetAura.Activate(sim)
			hunter.pet.ScentOfBloodPetAura.UpdateExpires(sim, sim.CurrentTime+time.Second*18)
		},
	})

	hunter.AddMajorCooldown(core.MajorCooldown{
		Spell: bwSpell,
		Type:  core.CooldownTypeDPS,
	})
}

func (hunter *Hunter) registerScentOfBlood() {
	if hunter.Talents.ScentOfBlood == 0 {
		return
	}

	actionID := core.ActionID{SpellID: 19605}

	hunter.pet.ScentOfBloodPetAura = hunter.pet.RegisterAura(core.Aura{
		Label:    "Scent of Blood Pet",
		ActionID: actionID,
		Duration: time.Second * 8,
	}).AttachMultiplicativePseudoStatBuff(&hunter.pet.PseudoStats.DamageDealtMultiplier, 1.4)

	core.MakePermanent(hunter.GetOrRegisterAura(core.Aura{
		Label:    "Scent of Blood",
		Duration: core.NeverExpires,
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.ProcMask.Matches(core.ProcMaskEmpty) || !(spell.ProcMask.Matches(core.ProcMaskRanged) || spell.ProcMask.Matches(core.ProcMaskMelee)) {
				return
			}

			if hunter.pet.ScentOfBloodPetAura.RemainingDuration(sim) < time.Second*8 && sim.Proc(0.05*float64(hunter.Talents.ScentOfBlood), "Scent of Blood") {
				hunter.pet.ScentOfBloodPetAura.Activate(sim)
			}
		},
	}))
}

func (hunter *Hunter) mortalShots() float64 {
	return 0.06 * float64(hunter.Talents.MortalShots)
}

func (hunter *Hunter) getUntamedTrapper() float64 {
	if hunter.Talents.UntamedTrapper {
		return 1
	}
	return 0
}

func (hunter *Hunter) applyTrapMastery() {
	if hunter.Talents.TrapMastery == 0 {
		return
	}

	hunter.OnSpellRegistered(func(spell *core.Spell) {
		if spell.Flags.Matches(SpellFlagTrap) {
			spell.BonusHitRating += 3.33 * float64(hunter.Talents.TrapMastery)
		}
	})

	hunter.OnSpellRegistered(func(spell *core.Spell) {
		if spell.Flags.Matches(SpellFlagTrap) {
			spell.DamageMultiplier *= 1 + 0.1*float64(hunter.Talents.TrapMastery)
		}
	})
}

func (hunter *Hunter) applyEfficiency() {
	hunter.OnSpellRegistered(func(spell *core.Spell) {
		// applies to Stings, Shots, and Volley
		if spell.Cost != nil && spell.Flags.Matches(SpellFlagSting|SpellFlagShot) || spell.SpellCode == SpellCode_HunterVolley {
			spell.Cost.Multiplier -= 2 * hunter.Talents.Efficiency
		}
	})
}
