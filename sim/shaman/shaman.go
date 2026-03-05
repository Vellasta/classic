package shaman

import (
	"time"

	"github.com/Vellasta/classic/sim/common/guardians"
	"github.com/Vellasta/classic/sim/core"
	"github.com/Vellasta/classic/sim/core/proto"
	"github.com/Vellasta/classic/sim/core/stats"
)

var TalentTreeSizes = [3]int{15, 16, 15}

const (
	SpellFlagShaman    = core.SpellFlagAgentReserved1
	SpellFlagTotem     = core.SpellFlagAgentReserved2
	SpellFlagLightning = core.SpellFlagAgentReserved3
)

func NewShaman(character *core.Character, talents string) *Shaman {
	shaman := &Shaman{
		Character: *character,
		Talents:   &proto.ShamanTalents{},
	}

	core.FillTalentsProto(shaman.Talents.ProtoReflect(), talents, TalentTreeSizes)
	shaman.EnableManaBar()

	// Add Shaman stat dependencies
	shaman.AddStatDependency(stats.Strength, stats.AttackPower, core.APPerStrength[character.Class])
	shaman.AddStatDependency(stats.Agility, stats.MeleeCrit, core.CritPerAgiAtLevel[character.Class]*core.CritRatingPerCritChance)
	shaman.AddStatDependency(stats.Agility, stats.Dodge, core.DodgePerAgiAtLevel[character.Class]*core.DodgeRatingPerDodgeChance)
	shaman.AddStatDependency(stats.Intellect, stats.SpellCrit, core.CritPerIntAtLevel[character.Class]*core.SpellCritRatingPerCritChance)
	shaman.AddStatDependency(stats.BonusArmor, stats.Armor, 1)
	shaman.PseudoStats.BlockValuePerStrength = .05 // 20 str = 1 block

	shaman.ApplyRockbiterImbue(shaman.getImbueProcMask(proto.WeaponImbue_RockbiterWeapon))
	shaman.ApplyFlametongueImbue(shaman.getImbueProcMask(proto.WeaponImbue_FlametongueWeapon))
	shaman.ApplyFrostbrandImbue(shaman.getImbueProcMask(proto.WeaponImbue_FrostbrandWeapon))
	shaman.ApplyWindfuryImbue(shaman.getImbueProcMask(proto.WeaponImbue_WindfuryWeapon))

	guardians.ConstructGuardians(&shaman.Character)

	return shaman
}

func (shaman *Shaman) getImbueProcMask(imbue proto.WeaponImbue) core.ProcMask {
	mask := core.ProcMaskUnknown
	if shaman.HasMHWeapon() && shaman.Consumes.MainHandImbue == imbue {
		mask |= core.ProcMaskMeleeMH
	}
	return mask
}

// Indexes into NextTotemDrops for self buffs
const (
	AirTotem int = iota
	EarthTotem
	FireTotem
	WaterTotem
)

const (
	SpellCode_ShamanNone int32 = iota

	SpellCode_ShamanChainHeal
	SpellCode_ShamanChainLightning
	SpellCode_ShamanEarthShock
	SpellCode_ShamanFireNovaTotem
	SpellCode_ShamanFlameShock
	SpellCode_ShamanFrostShock
	SpellCode_ShamanHealingWave
	SpellCode_ShamanLesserHealingWave
	SpellCode_ShamanLightningBolt
	SpellCode_ShamanLightningShield
	SpellCode_ShamanMagmaTotem
	SpellCode_ShamanSearingTotem
	SpellCode_ShamanStormstrike
)

// Shaman represents a shaman character.
type Shaman struct {
	core.Character

	Talents *proto.ShamanTalents

	// Spells
	ChainHeal            []*core.Spell
	ChainLightning       []*core.Spell
	EarthShield          *core.Spell
	EarthShock           []*core.Spell
	ElementalMastery     *core.Spell
	FireNovaTotem        []*core.Spell
	FlameShock           []*core.Spell
	FrostShock           []*core.Spell
	GraceOfAirTotem      []*core.Spell
	HealingStreamTotem   []*core.Spell
	HealingWave          []*core.Spell
	LesserHealingWave    []*core.Spell
	LightningBolt        []*core.Spell
	LightningShield      []*core.Spell
	LightningShieldProcs []*core.Spell // The damage component of lightning shield is a separate spell
	MagmaTotem           []*core.Spell
	ManaSpringTotem      []*core.Spell
	SearingTotem         []*core.Spell
	StoneskinTotem       []*core.Spell
	Stormstrike          *core.Spell
	StrengthOfEarthTotem []*core.Spell
	TremorTotem          *core.Spell
	WindfuryTotem        []*core.Spell
	WindfuryWeaponMH     *core.Spell
	WindfuryWeaponOH     *core.Spell
	WindwallTotem        []*core.Spell

	// Auras
	ClearcastingAura     *core.Aura
	LightningShieldAuras []*core.Aura

	// Totems
	ActiveTotems     [4]*core.Spell
	ActiveTotemBuffs [4]*core.Aura
	TotemExpirations [4]time.Duration // The expiration time of each totem (earth, air, fire, water).

	EarthTotems []*core.Spell
	FireTotems  []*core.Spell
	WaterTotems []*core.Spell
	AirTotems   []*core.Spell
	Totems      *proto.ShamanTotems

	WindfuryTotemPeriodicActions      []*core.PendingAction
	ActiveWindfuryTotemPeriodicAction *core.PendingAction

	// Shield
	ActiveShield     *core.Spell // Tracks the Shaman's active shield spell
	ActiveShieldAura *core.Aura

	ChainLightningBounceCoefficient float64
}

// Implemented by each Shaman spec.
type ShamanAgent interface {
	core.Agent

	// The Shaman controlled by this Agent.
	GetShaman() *Shaman
}

func (shaman *Shaman) GetCharacter() *core.Character {
	return &shaman.Character
}

func (shaman *Shaman) AddRaidBuffs(_ *proto.RaidBuffs) {
	// Buffs are handled explicitly through APLs now
}

func (shaman *Shaman) Initialize() {
	// Core abilities
	shaman.registerChainLightningSpell()
	shaman.registerLightningBoltSpell()
	shaman.registerLightningShieldSpell()
	shaman.registerShocks()
	shaman.registerStormstrikeSpell()

	// Imbues
	// In the Initialize due to frost brand adding the aura to the enemy
	shaman.RegisterRockbiterImbue(shaman.getImbueProcMask(proto.WeaponImbue_RockbiterWeapon))
	shaman.RegisterFlametongueImbue(shaman.getImbueProcMask(proto.WeaponImbue_FlametongueWeapon))
	shaman.RegisterWindfuryImbue(shaman.getImbueProcMask(proto.WeaponImbue_WindfuryWeapon))
	shaman.RegisterFrostbrandImbue(shaman.getImbueProcMask(proto.WeaponImbue_FrostbrandWeapon))

	// Totems
	shaman.registerStrengthOfEarthTotemSpell()
	shaman.registerStoneskinTotemSpell()
	shaman.registerTremorTotemSpell()
	shaman.registerSearingTotemSpell()
	shaman.registerMagmaTotemSpell()
	shaman.registerFireNovaTotemSpell()
	shaman.registerHealingStreamTotemSpell()
	shaman.registerManaSpringTotemSpell()
	shaman.registerWindfuryTotemSpell()
	shaman.registerGraceOfAirTotemSpell()
	shaman.registerWindwallTotemSpell()
}

func (shaman *Shaman) Reset(_ *core.Simulation) {
	shaman.ActiveShield = nil
	shaman.ActiveShieldAura = nil

	for i := range []int{EarthTotem, FireTotem, WaterTotem, AirTotem} {
		shaman.ActiveTotems[i] = nil
		shaman.TotemExpirations[i] = 0
		shaman.ActiveTotemBuffs[i] = nil
	}
}
