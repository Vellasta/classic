package hunter

import (
	"time"

	"github.com/Vellasta/classic/sim/core"
)

func (hunter *Hunter) getKillCommandConfig(rank int) core.SpellConfig {
	spellId := [2]int32{0, 41827}[rank]
	level := [2]int{0, 40}[rank]

	spellConfig := core.SpellConfig{
		SpellCode:     SpellCode_HunterKillCommand,
		ActionID:      core.ActionID{SpellID: spellId},
		SpellSchool:   core.SpellSchoolPhysical,
		DefenseType:   core.DefenseTypeMelee,
		ProcMask:      core.ProcMaskMeleeSpecial,
		Flags:         core.SpellFlagMeleeMetrics | core.SpellFlagAPL,
		Rank:          rank,
		RequiredLevel: level,

		ManaCost: core.ManaCostOptions{
			FlatCost: 0.05 * hunter.BaseMana,
		},

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    hunter.NewTimer(),
				Duration: time.Second * 8,
			},
		},

		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return hunter.KillCommandAura.IsActive() && hunter.pet.IsActive()
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			hunter.pet.killCommand.Cast(sim, target)
		},
	}

	return spellConfig
}

func (hp *HunterPet) newKillCommand() *core.Spell {
	return hp.RegisterSpell(core.SpellConfig{
		SpellCode:   SpellCode_HunterPetKillCommand,
		ActionID:    core.ActionID{SpellID: 41827},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagMeleeMetrics,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			damage := 0.8 * spell.MeleeAttackPower(target)
			spell.CalcAndDealDamage(sim, target, damage, spell.OutcomeMeleeSpecialHitAndCrit)
		},
	})
}

func (hunter *Hunter) registerKillCommandSpell() {
	if !hunter.Talents.KillCommand {
		return
	}

	hunter.KillCommandAura = hunter.RegisterAura(core.Aura{
		Label:    "Kill Command Aura",
		ActionID: core.ActionID{SpellID: 41827},
		Duration: time.Second * 4,
	})

	hunter.RegisterAura(core.Aura{
		Label:    "Kill Command Trigger",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if result.DidCrit() {
				hunter.KillCommandAura.Activate(sim)
			}
		},
	})

	rank := map[int32]int{
		25: 1,
		40: 1,
		50: 1,
		60: 1,
	}[hunter.Level]

	config := hunter.getKillCommandConfig(rank)
	hunter.KillCommand = hunter.GetOrRegisterSpell(config)
}
