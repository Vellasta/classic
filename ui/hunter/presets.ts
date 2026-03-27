import { Phase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	AgilityElixir,
	Alcohol,
	AttackPowerBuff,
	Conjured,
	Consumes,
	Debuffs,
	Flask,
	Food,
	HealthElixir,
	IndividualBuffs,
	ManaRegenElixir,
	Potions,
	Profession,
	Race,
	RaidBuffs,
	SapperExplosive,
	SaygesFortune,
	SpellPowerBuff,
	StrengthBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
	BlastedLandsBuff,
} from '../core/proto/common.js';
import {
	Hunter_Options as HunterOptions,
	Hunter_Options_Ammo as Ammo,
	Hunter_Options_PetAttackSpeed as PetAttackSpeed,
	Hunter_Options_PetType as PetType,
	Hunter_Options_QuiverBonus,
} from '../core/proto/hunter.js';
import { SavedTalents } from '../core/proto/ui.js';
import BMAPL from './apls/bm.apl.json';
import MMOldAPL from './apls/mm.old.apl.json';
import MMAPL from './apls/mm.apl.json';
import MMAPLBasic from './apls/mm.basic.apl.json';
import MMAPLAdvanced from './apls/mm.advanced.apl.json';
import SVAPL from './apls/sv.apl.json';
import PreraidBMGear from './gear_sets/preraid.bm.gear.json';
import PreraidMMGear from './gear_sets/preraid.mm.gear.json';
import PreraidSVGear from './gear_sets/preraid.sv.gear.json';
import MCBMGear from './gear_sets/mc.bm.gear.json';
import MCMMGear from './gear_sets/mc.mm.gear.json';
import MCSVGear from './gear_sets/mc.sv.gear.json';
import BWLBMGear from './gear_sets/bwl.bm.gear.json';
import BWLMMGear from './gear_sets/bwl.mm.gear.json';
import BWLSVGear from './gear_sets/bwl.sv.gear.json';
import AQ40BMGear from './gear_sets/aq40.bm.gear.json';
import AQ40MMGear from './gear_sets/aq40.mm.gear.json';
import AQ40SVGear from './gear_sets/aq40.sv.gear.json';
import NaxxBMGear from './gear_sets/naxx.bm.gear.json';
import NaxxMMGear from './gear_sets/naxx.mm.gear.json';
import NaxxSVGear from './gear_sets/naxx.sv.gear.json';
import Kara40BMGear from './gear_sets/kara40.bm.gear.json';
import Kara40MMGear from './gear_sets/kara40.mm.gear.json';
import Kara40SVGear from './gear_sets/kara40.sv.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.
///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearPreraidBM = PresetUtils.makePresetGear('Pre-Raid BM', PreraidBMGear, { talentTree: 0 });
export const GearPreraidMM = PresetUtils.makePresetGear('Pre-Raid MM', PreraidMMGear, { talentTree: 1 });
export const GearPreraidSV = PresetUtils.makePresetGear('Pre-Raid SV', PreraidSVGear, { talentTree: 2 });
export const GearMCBM = PresetUtils.makePresetGear('MC BM', MCBMGear, { talentTree: 0 });
export const GearMCMM = PresetUtils.makePresetGear('MC MM', MCMMGear, { talentTree: 1 });
export const GearMCSV = PresetUtils.makePresetGear('MC SV', MCSVGear, { talentTree: 2 });
export const GearBWLBM = PresetUtils.makePresetGear('BWL BM', BWLBMGear, { talentTree: 0 });
export const GearBWLMM = PresetUtils.makePresetGear('BWL MM', BWLMMGear, { talentTree: 1 });
export const GearBWLSV = PresetUtils.makePresetGear('BWL SV', BWLSVGear, { talentTree: 2 });
export const GearAQ40BM = PresetUtils.makePresetGear('AQ40 BM', AQ40BMGear, { talentTree: 0 });
export const GearAQ40MM = PresetUtils.makePresetGear('AQ40 MM', AQ40MMGear, { talentTree: 1 });
export const GearAQ40SV = PresetUtils.makePresetGear('AQ40 SV', AQ40SVGear, { talentTree: 2 });
export const GearNaxxBM = PresetUtils.makePresetGear('Naxx BM', NaxxBMGear, { talentTree: 0 });
export const GearNaxxMM = PresetUtils.makePresetGear('Naxx MM', NaxxMMGear, { talentTree: 1 });
export const GearNaxxSV = PresetUtils.makePresetGear('Naxx SV', NaxxSVGear, { talentTree: 2 });
export const GearKara40BM = PresetUtils.makePresetGear('Kara40 BM', Kara40BMGear, { talentTree: 0 });
export const GearKara40MM = PresetUtils.makePresetGear('Kara40 MM', Kara40MMGear, { talentTree: 1 });
export const GearKara40SV = PresetUtils.makePresetGear('Kara40 SV', Kara40SVGear, { talentTree: 2 });

export const GearPresets = {
	[Phase.Phase1]: [
		GearPreraidBM,
		GearPreraidMM, 
		GearPreraidSV, 
		GearMCBM,
		GearMCMM, 
		GearMCSV, 
		GearBWLBM,
		GearBWLMM, 
		GearBWLSV, 
		GearAQ40BM,
		GearAQ40MM, 
		GearAQ40SV, 
		GearNaxxBM,
		GearNaxxMM, 
		GearNaxxSV, 
		GearKara40BM,
		GearKara40MM, 
		GearKara40SV
	],
};

export const DefaultGear = GearKara40MM;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLBM = PresetUtils.makePresetAPLRotation('Beast Mastery', BMAPL, { talentTree: 0 });
export const APLMM = PresetUtils.makePresetAPLRotation('Marksmanship (Standard)', MMAPL, { talentTree: 1 });
export const APLMMBasic = PresetUtils.makePresetAPLRotation('Marksmanship (Basic)', MMAPLBasic, { talentTree: 1 });
export const APLMMAdvanced = PresetUtils.makePresetAPLRotation('Marksmanship (Advanced)', MMAPLAdvanced, { talentTree: 1 });
export const APLMMOld = PresetUtils.makePresetAPLRotation('Marksmanship (Old)', MMOldAPL, { talentTree: 1 });
export const APLSV = PresetUtils.makePresetAPLRotation('Survival', SVAPL, { talentTree: 2 });

export const APLPresets = {
	[Phase.Phase1]: [APLBM, APLMMBasic, APLMM, APLMMAdvanced, APLSV, APLMMOld],

};

export const DefaultAPL = APLPresets[Phase.Phase1][1];

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsBM = PresetUtils.makePresetTalents('Beast Mastery', SavedTalents.create({ talentsString: '5500000150053102221-050520322' }));
export const TalentsMM = PresetUtils.makePresetTalents('Marksmanship', SavedTalents.create({ talentsString: '550000012-0525210250123251-002' }));
export const TalentsSV = PresetUtils.makePresetTalents('Survival', SavedTalents.create({ talentsString: '550000015-0000000000000000-35202000111212331251' }));

export const TalentPresets = {
	[Phase.Phase1]: [TalentsBM, TalentsMM, TalentsSV],
};

export const DefaultTalents = TalentPresets[Phase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = HunterOptions.create({
	ammo: Ammo.Doomshot,
	quiverBonus: Hunter_Options_QuiverBonus.Speed15,
	petAttackSpeed: PetAttackSpeed.OneTwo,
	petType: PetType.Cat,
	petUptime: 1,
});

export const DefaultConsumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	alcohol: Alcohol.AlcoholRumseyRumBlackLabel,
	attackPowerBuff: AttackPowerBuff.JujuMight,
	defaultConjured: Conjured.ConjuredDemonicRune,
	defaultPotion: Potions.MajorManaPotion,
	dragonBreathChili: true,
	flask: Flask.FlaskOfSupremePower,
	food: Food.FoodGrilledSquid,
	healthElixir: HealthElixir.ElixirOfFortitude,
	mainHandImbue: WeaponImbue.ElementalSharpeningStone,
	manaRegenElixir: ManaRegenElixir.MagebloodPotion,
	offHandImbue: WeaponImbue.ElementalSharpeningStone,
	petAttackPowerConsumable: 1,
	petAgilityConsumable: 1,
	petStrengthConsumable: 1,
	sapperExplosive: SapperExplosive.SapperUnknown,
	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.SpiritOfZanza,
	blastedLandsBuff: BlastedLandsBuff.GroundScorpokAssay
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	battleShout: TristateEffect.TristateEffectImproved,
	divineSpirit: true,
	fireResistanceAura: true,
	fireResistanceTotem: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	powerWordFortitude: TristateEffect.TristateEffectImproved,
	graceOfAirTotem: TristateEffect.TristateEffectMissing,
	trueshotAura: true,
	leaderOfThePack: true,
	manaSpringTotem: TristateEffect.TristateEffectRegular,
	strengthOfEarthTotem: TristateEffect.TristateEffectImproved,
	windfuryTotem: true,
	flametongueTotem: true
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfMight: TristateEffect.TristateEffectImproved,
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
	fengusFerocity: false,
	moldarsMoxie: false,
	rallyingCryOfTheDragonslayer: false,
	saygesFortune: SaygesFortune.SaygesUnknown,
	slipkiksSavvy: false,
	songflowerSerenade: false,
	spiritOfZandalar: false,
	warchiefsBlessing: false,
});

export const DefaultDebuffs = Debuffs.create({
	curseOfRecklessness: true,
	exposeArmor: TristateEffect.TristateEffectImproved,
	faerieFire: true,
	huntersMark: true,
	improvedScorch: true,
	judgementOfWisdom: true,
	stormstrike: false,
	sunderArmor: true,
});

export const OtherDefaults = {
	distanceFromTarget: 12,
	weaponSkillBook: false,
	profession1: Profession.Enchanting,
	profession2: Profession.Engineering,
	race: Race.RaceTroll,
};
