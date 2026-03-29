import * as BuffDebuffInputs from '../core/components/inputs/buffs_debuffs';
import * as ConsumablesInputs from '../core/components/inputs/consumables.js';
import * as OtherInputs from '../core/components/other_inputs.js';
import { Phase } from '../core/constants/other.js';
import { IndividualSimUI, registerSpecConfig } from '../core/individual_sim_ui.js';
import { Player } from '../core/player.js';
import { PartyBuffs, PseudoStat, Spec, Stat } from '../core/proto/common.js';
import { Stats } from '../core/proto_utils/stats.js';
import * as HunterInputs from './inputs.js';
import * as Presets from './presets.js';

const SPEC_CONFIG = registerSpecConfig(Spec.SpecHunter, {
	cssClass: 'hunter-sim-ui',
	cssScheme: 'hunter',
	// List any known bugs / issues here and they'll be shown on the site.
	knownIssues: [],
	warnings: [],

	// All stats for which EP should be calculated.
	epStats: [
		// Attributes
		Stat.StatStrength,
		Stat.StatAgility,
		// Physical
		Stat.StatAttackPower,
		Stat.StatRangedAttackPower,
		Stat.StatMeleeHit,
		Stat.StatMeleeCrit,
		Stat.StatArmorPenetration,
		Stat.StatMeleeHaste,
		// Spell
		Stat.StatSpellPower,
	],
	epPseudoStats: [
		PseudoStat.PseudoStatMainHandDps,
		PseudoStat.PseudoStatOffHandDps,
		PseudoStat.PseudoStatRangedDps,
	],
	// Reference stat against which to calculate EP.
	epReferenceStat: Stat.StatRangedAttackPower,
	// Which stats to display in the Character Stats section, at the bottom of the left-hand sidebar.
	displayStats: [
		// Primary
		Stat.StatMana,
		// Attributes
		Stat.StatStrength,
		Stat.StatAgility,
		// Physical
		Stat.StatAttackPower,
		Stat.StatRangedAttackPower,
		Stat.StatMeleeHit,
		Stat.StatMeleeCrit,
		Stat.StatArmorPenetration,
		Stat.StatMeleeHaste,
		// Spell
		Stat.StatSpellDamage,
		Stat.StatNaturePower,
		Stat.StatArcanePower,
		Stat.StatMP5,
	],
	displayPseudoStats: [PseudoStat.PseudoStatMeleeSpeedMultiplier, PseudoStat.PseudoStatRangedSpeedMultiplier],

	defaults: {
		race: Presets.OtherDefaults.race,
		// Default equipped gear.
		gear: Presets.DefaultGear.gear,
		// Default EP weights for sorting gear in the gear picker.
		epWeights: Stats.fromMap(
			{
				[Stat.StatStrength]: 1.0,
				[Stat.StatAgility]: 2.8,
				[Stat.StatStamina]: 0.02,
				[Stat.StatIntellect]: 0.02,
				[Stat.StatAttackPower]: 0.5,
				[Stat.StatRangedAttackPower]: 0.5,
				[Stat.StatMeleeHit]: 20.0,
				[Stat.StatMeleeCrit]: 33.0,
				[Stat.StatArmorPenetration]: 0.3,
				[Stat.StatMeleeHaste]: 20.0,
				[Stat.StatSpellPower]: 0.03,
				[Stat.StatNaturePower]: 0.01,
				[Stat.StatArcanePower]: 0.01,
				[Stat.StatMP5]: 0.05,
			},
			{
				[PseudoStat.PseudoStatMainHandDps]: 10.0,
				[PseudoStat.PseudoStatOffHandDps]: 10.0,
				[PseudoStat.PseudoStatRangedDps]: 14.0,
				[PseudoStat.PseudoStatUnarmedSkill]: 2.0,
				[PseudoStat.PseudoStatDaggersSkill]: 2.0,
				[PseudoStat.PseudoStatSwordsSkill]: 2.0,
				[PseudoStat.PseudoStatMacesSkill]: 2.0,
				[PseudoStat.PseudoStatAxesSkill]: 2.0,
				[PseudoStat.PseudoStatTwoHandedSwordsSkill]: 2.0,
				[PseudoStat.PseudoStatTwoHandedMacesSkill]: 2.0,
				[PseudoStat.PseudoStatTwoHandedAxesSkill]: 2.0,
				[PseudoStat.PseudoStatPolearmsSkill]: 2.0,
				[PseudoStat.PseudoStatStavesSkill]: 2.0,
			},
		),
		// Default consumes settings.
		consumes: Presets.DefaultConsumes,
		// Default talents.
		talents: Presets.DefaultTalents.data,
		// Default spec-specific settings.
		specOptions: Presets.DefaultOptions,
		other: Presets.OtherDefaults,
		// Default raid/party buffs settings.
		raidBuffs: Presets.DefaultRaidBuffs,
		partyBuffs: PartyBuffs.create({}),
		individualBuffs: Presets.DefaultIndividualBuffs,
		debuffs: Presets.DefaultDebuffs,
	},

	// IconInputs to include in the 'Player' section on the settings tab.
	playerIconInputs: [HunterInputs.PetTypeInput, HunterInputs.WeaponAmmo, HunterInputs.QuiverInput],
	// Inputs to include in the 'Rotation' section on the settings tab.
	rotationInputs: HunterInputs.HunterRotationConfig,
	petConsumeInputs: [ConsumablesInputs.PetAttackPowerConsumable, ConsumablesInputs.PetAgilityConsumable, ConsumablesInputs.PetStrengthConsumable],
	// Buff and Debuff inputs to include/exclude, overriding the EP-based defaults.
	includeBuffDebuffInputs: [ConsumablesInputs.DragonBreathChili, BuffDebuffInputs.SpellScorchDebuff, BuffDebuffInputs.StaminaBuff],
	excludeBuffDebuffInputs: [],
	// Inputs to include in the 'Other' section on the settings tab.
	otherInputs: {
		inputs: [
			//HunterInputs.NewRaptorStrike,
			HunterInputs.PetAttackSpeedInput,
			HunterInputs.PetUptime,
			OtherInputs.DistanceFromTarget,
			OtherInputs.TankAssignment,
			OtherInputs.InFrontOfTarget,
			OtherInputs.weaponSkillBook,
		],
	},
	encounterPicker: {
		// Whether to include 'Execute Duration (%)' in the 'Encounter' section of the settings tab.
		showExecuteProportion: false,
	},

	presets: {
		// Preset talents that the user can quickly select.
		talents: [...Presets.TalentPresets[Phase.Phase1]],
		// Preset rotations that the user can quickly select.
		rotations: [...Presets.APLPresets[Phase.Phase1]],
		// Preset gear configurations that the user can quickly select.
		gear: [...Presets.GearPresets[Phase.Phase1]],
	},

	autoRotation: player => {
		const isMelee = false;
		//player.hasRune(ItemSlot.ItemSlotWaist, HunterRune.RuneBeltMeleeSpecialist) ||
		//player.hasRune(ItemSlot.ItemSlotFeet, HunterRune.RuneBootsDualWieldSpecialization) ||
		//player.hasRune(ItemSlot.ItemSlotFeet, HunterRune.RuneBootsWyvernStrike);

		const talentTree = player.getTalentTree();
		if (talentTree === 0) {
			return Presets.APLBM.rotation.rotation!;
		} else if (talentTree === 1) {
			return Presets.APLMM.rotation.rotation!;
		} else if (talentTree === 2) {
			return Presets.APLSV.rotation.rotation!;
		}

		return Presets.DefaultAPL.rotation.rotation!;

		// COMMENTING OUT TO SAVE FOR FUTURE IMPLEMENTATION.
		// if (isMelee) {
		// 	switch (level) {
		// 		case 25:
		// 			return Presets.APLMeleeWeavePhase1.rotation.rotation!;
		// 		case 40:
		// 			return Presets.APLMeleePhase2.rotation.rotation!;
		// 		case 50:
		// 			return Presets.APLMeleeBmPhase3.rotation.rotation!;
		// 		case 60:
		// 			return Presets.APLWeavePhase4.rotation.rotation!;
		// 	}
		// } else {
		// 	switch (level) {
		// 		case 25:
		// 			return Presets.APLMeleeWeavePhase1.rotation.rotation!;
		// 		case 40:
		// 			return player.getTalentTree() === 1 ? Presets.APLRangedMmPhase2.rotation.rotation! : Presets.APLRangedBmPhase2.rotation.rotation!;
		// 		case 50:
		// 			return Presets.APLRangedMmPhase3.rotation.rotation!;
		// 		case 60:
		// 			return Presets.APLRangedPhase4.rotation.rotation!;
		// 	}
		// }
		//throw new Error('Auto rotation not supported for your current configuration.');
	},

	raidSimPresets: [
		// Raid sim presets dont work very well with SoD specs between phases
		// and we dont support raid sim atm so just comment this out
		// {
		// 	spec: Spec.SpecHunter,
		// 	tooltip: 'Beast Mastery Hunter',
		// 	defaultName: 'Beast Mastery',
		// 	iconUrl: getSpecIcon(Class.ClassHunter, 0),
		// 	talents: Presets.DefaultTalentsBeastMastery.data,
		// 	specOptions: Presets.BMDefaultOptions,
		// 	consumes: Presets.DefaultConsumes,
		// 	defaultFactionRaces: {
		// 		[Faction.Unknown]: Race.RaceUnknown,
		// 		[Faction.Alliance]: Race.RaceNightElf,
		// 		[Faction.Horde]: Race.RaceOrc,
		// 	},
		// 	defaultGear: {
		// 		[Faction.Unknown]: {},
		// 		[Faction.Alliance]: {
		// 			1: Presets.GearPresets[Phase.Phase1][0].gear,
		// 		},
		// 		[Faction.Horde]: {
		// 			1: Presets.GearPresets[Phase.Phase1][0].gear,
		// 		},
		// 	},
		// },
		// {
		// 	spec: Spec.SpecHunter,
		// 	tooltip: 'Marksmanship Hunter',
		// 	defaultName: 'Marksmanship',
		// 	iconUrl: getSpecIcon(Class.ClassHunter, 1),
		// 	talents: Presets.DefaultTalentsMarksman.data,
		// 	specOptions: Presets.DefaultOptions,
		// 	consumes: Presets.DefaultConsumes,
		// 	defaultFactionRaces: {
		// 		[Faction.Unknown]: Race.RaceUnknown,
		// 		[Faction.Alliance]: Race.RaceNightElf,
		// 		[Faction.Horde]: Race.RaceOrc,
		// 	},
		// 	defaultGear: {
		// 		[Faction.Unknown]: {},
		// 		[Faction.Alliance]: {
		// 			1: Presets.GearPresets[Phase.Phase1][1].gear,
		// 		},
		// 		[Faction.Horde]: {
		// 			1: Presets.GearPresets[Phase.Phase1][1].gear,
		// 		},
		// 	},
		// },
		// {
		// 	spec: Spec.SpecHunter,
		// 	tooltip: 'Survival Hunter',
		// 	defaultName: 'Survival',
		// 	iconUrl: getSpecIcon(Class.ClassHunter, 2),
		// 	talents: Presets.DefaultTalentsSurvival.data,
		// 	specOptions: Presets.DefaultOptions,
		// 	consumes: Presets.DefaultConsumes,
		// 	defaultFactionRaces: {
		// 		[Faction.Unknown]: Race.RaceUnknown,
		// 		[Faction.Alliance]: Race.RaceNightElf,
		// 		[Faction.Horde]: Race.RaceOrc,
		// 	},
		// 	defaultGear: {
		// 		[Faction.Unknown]: {},
		// 		[Faction.Alliance]: {
		// 			1: Presets.GearPresets[Phase.Phase1][2].gear,
		// 		},
		// 		[Faction.Horde]: {
		// 			1: Presets.GearPresets[Phase.Phase1][2].gear,
		// 		},
		// 	},
		// },
	],
});

export class HunterSimUI extends IndividualSimUI<Spec.SpecHunter> {
	constructor(parentElem: HTMLElement, player: Player<Spec.SpecHunter>) {
		super(parentElem, player, SPEC_CONFIG);
	}
}
