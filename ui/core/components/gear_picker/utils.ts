import { ItemSlot } from '../../proto/common';
import ItemList, { GearData, ItemData, ItemListType } from './item_list';
import { DatabaseFilters, RepSource, UIEnchant, UIFaction, UIItem, UIItem_FactionRestriction } from '../../proto/ui';
import { Stat } from '../../proto/common';

const emptySlotIcons: Record<ItemSlot, string> = {
	[ItemSlot.ItemSlotHead]: '/classic/assets/item_slots/head.jpg',
	[ItemSlot.ItemSlotNeck]: '/classic/assets/item_slots/neck.jpg',
	[ItemSlot.ItemSlotShoulder]: '/classic/assets/item_slots/shoulders.jpg',
	[ItemSlot.ItemSlotBack]: '/classic/assets/item_slots/shirt.jpg',
	[ItemSlot.ItemSlotChest]: '/classic/assets/item_slots/chest.jpg',
	[ItemSlot.ItemSlotWrist]: '/classic/assets/item_slots/wrists.jpg',
	[ItemSlot.ItemSlotHands]: '/classic/assets/item_slots/hands.jpg',
	[ItemSlot.ItemSlotWaist]: '/classic/assets/item_slots/waist.jpg',
	[ItemSlot.ItemSlotLegs]: '/classic/assets/item_slots/legs.jpg',
	[ItemSlot.ItemSlotFeet]: '/classic/assets/item_slots/feet.jpg',
	[ItemSlot.ItemSlotFinger1]: '/classic/assets/item_slots/finger.jpg',
	[ItemSlot.ItemSlotFinger2]: '/classic/assets/item_slots/finger.jpg',
	[ItemSlot.ItemSlotTrinket1]: '/classic/assets/item_slots/trinket.jpg',
	[ItemSlot.ItemSlotTrinket2]: '/classic/assets/item_slots/trinket.jpg',
	[ItemSlot.ItemSlotMainHand]: '/classic/assets/item_slots/mainhand.jpg',
	[ItemSlot.ItemSlotOffHand]: '/classic/assets/item_slots/offhand.jpg',
	[ItemSlot.ItemSlotRanged]: '/classic/assets/item_slots/ranged.jpg',
};
export function getEmptySlotIconUrl(slot: ItemSlot): string {
	return emptySlotIcons[slot];
}

const indexToStatFunctionMap: Record<number, Function> = {
	[0]: getStrengthTooltip,
	[1]: getAgilityTooltip,
	[2]: getStaminaTooltip,
	[3]: getIntellectTooltip,
	[4]: getSpiritTooltip,
	[5]: getSpellPowerTooltip,
	[6]: getArcanePowerTooltip,
	[7]: getFirePowerTooltip,
	[8]: getFrostPowerTooltip,
	[9]: getHolyPowerTooltip,
	[10]: getNaturePowerTooltip,
	[11]: getShadowPowerTooltip,
	[12]: getMP5Tooltip,
	[13]: getSpellHitTooltip,
	[14]: getSpellCritTooltip,
	[15]: getSpellHasteTooltip,
	[16]: getSpellPenetrationTooltip,
	[17]: getAttackPowerTooltip,
	[18]: getMeleeHitTooltip,
	[19]: getMeleeCritTooltip,
	[20]: getMeleeHasteTooltip,
	[21]: getArmorPenetrationTooltip,
	[27]: getRangedAttackPowerTooltip,
	[28]: getDefenseTooltip,
	[29]: getBlockTooltip,
	[30]: getBlockValueTooltip,
	[31]: getDodgeTooltip,
	[32]: getParryTooltip,
	[35]: getArcaneResistanceTooltip,
	[36]: getFireResistanceTooltip,
	[37]: getFrostResistanceTooltip,
	[38]: getNatureResistanceTooltip,
	[39]: getShadowResistanceTooltip,
	[41]: getHealingPowerTooltip,
	[42]: getSpellDamageTooltip,
	[43]: getFeralAttackPowerTooltip,
}

const qualityToColorMap: Record<number, String> = {
	[0]: "#9d9d9d",
	[1]: "#ffffff",
	[2]: "#1eff00",
	[3]: "#0070dd",
	[4]: "#a335ee",
	[5]: "#ff8000",
}

const itemTypeToStringMap: Record<number, String> = {
	[1]: "Head",
	[2]: "Neck",
	[3]: "Shoulder",
	[4]: "Back",
	[5]: "Chest",
	[6]: "Wrist",
	[7]: "Hands",
	[8]: "Waist",
	[9]: "Legs",
	[10]: "Feet",
	[11]: "Finger",
	[12]: "Trinket",
	[13]: "Weapon",
	[14]: "Ranged",
}

const armorTypeToStringMap: Record<number, String> = {
	[0]: "Cloth",
	[1]: "Cloth",
	[2]: "Leather",
	[3]: "Mail",
	[4]: "Plate",
}

const handTypeToStringMap: Record<number, String> = {
	[1]: "Main-Hand",
	[2]: "One-Hand",
	[3]: "Off-Hand",
	[4]: "Two-Hand",
}

const meleeWeaponTypeToStringMap: Record<number, String> = {
	[1]: "Axe",
	[2]: "Dagger",
	[3]: "Fist",
	[4]: "Mace",
	[5]: "Off-Hand",
	[6]: "Polearm",
	[7]: "Shield",
	[8]: "Staff",
	[9]: "Sword",
}

const rangedWeaponTypeToStringMap: Record<number, String> = {
	[1]: "Bow",
	[2]: "Crossbow",
	[3]: "Gun",
	[4]: "Idol",
	[5]: "Libram",
	[6]: "Thrown",
	[7]: "Totem",
	[8]: "Wand",
	[9]: "Sigil",
}

const weaponSkillToStringFunctionMap: Record<number, Function> = {
	[1]: getAxesWeaponSkillTooltip,
	[2]: getSwordsWeaponSkillTooltip,
	[3]: getMacesWeaponSkillTooltip,
	[4]: getDaggersWeaponSkillTooltip,
	[5]: getUnarmedWeaponSkillTooltip,
	[6]: getTwoHandedAxesWeaponSkillTooltip,
	[7]: getTwoHandedSwordsWeaponSkillTooltip,
	[8]: getTwoHandedMacesWeaponSkillTooltip,
	[9]: getPolearmsWeaponSkillTooltip,
	[10]: getStavesWeaponSkillTooltip,
	[11]: getThrownWeaponSkillTooltip,
	[12]: getBowsWeaponSkillTooltip,
	[13]: getCrossbowsWeaponSkillTooltip,
	[14]: getGunsWeaponSkillTooltip,
}


export function createItemTooltip(item: ItemListType): string {
	const nameTooltip = getNameTooltip(item)
	const ilvlTooltip = getIlvlTooltip(item)
	const itemTypeTooltip = getitemTypeTooltip(item)
	const weaponTooltip = getWeaponTooltip(item)
	const armorTooltip = getArmorTooltip(item.stats)
	const baseStatsTooltip = getBaseStatsTooltip(item.stats)
	const bonusStatsTooltip = getBonusStatsTooltip(item.stats)
	const weaponSkillTooltip = getWeaponSkillTooltip(item)

	const tooltipList: string[] = [
		nameTooltip, 
		ilvlTooltip, 
		itemTypeTooltip,
		weaponTooltip, 
		armorTooltip, 
		baseStatsTooltip, 
		bonusStatsTooltip,
		weaponSkillTooltip
	]
	const filteredTooltipList = tooltipList.filter((s) => s != "")
	return filteredTooltipList.join('<br>')
}

function getNameTooltip(item: any): string {
	return `<span style="color: ${qualityToColorMap[item.quality]}; font-size: 112.5%;">${item.name} </span>`
}

function getIlvlTooltip(item: any): string {
	return `<span style="color: gold;">Item Level ${item.ilvl}</span>`
}

function getitemTypeTooltip(item: any): string {
	const itemType = itemTypeToStringMap[item.type]
	if (item.type <= 10) {
		const armorType = armorTypeToStringMap[item.armorType]
		return `<div style="overflow: hidden; display: inline; gap: 20px;"><span style="float: left; margin-right: 40px;">${itemType}</span><span style="float: right;">${armorType}</span></div>`
	}
	if (item.type == 13) {
		const meleeWeaponType = meleeWeaponTypeToStringMap[item.weaponType]
		const handType = handTypeToStringMap[item.handType]
		return `<div style="overflow: hidden; display: inline; gap: 20px;"><span style="float: left; margin-right: 40px;">${handType}</span><span style="float: right;">${meleeWeaponType}</span></div>`
	}
	if (item.type == 14) {
		const rangedWeaponType = rangedWeaponTypeToStringMap[item.rangedWeaponType]
		return `<div style="overflow: hidden; display: inline; gap: 20px;"><span style="float: left; margin-right: 40px;">Ranged</span><span style="float: right;">${rangedWeaponType}</span></div>`
	}

	return `${itemType}`
}

function getWeaponTooltip(item: any): string {
	if (item.weaponDamageMin > 0) {
		const dmg = `${item.weaponDamageMin} - ${item.weaponDamageMax} Damage`
		const speed = `Speed ${item.weaponSpeed.toFixed(2)}`
		const dps = `(${((item.weaponDamageMin + item.weaponDamageMax) / (2 * item.weaponSpeed)).toFixed(2)} damage per second)`
		return `<div style="overflow: hidden; display: inline;"><span style="float: left; margin-right: 40px;">${dmg}</span><span style="float: right;">${speed}</span></div><br>${dps}`
	}
	return ""
}

function getArmorTooltip(stats: any): string {
	if (stats[26] != 0) {
		return `${stats[26] + stats[40]} Armor`
	}
	return ""
}

function getBaseStatsTooltip(stats: any): string {
	const tooltipList: String[] = []
	for (let i = 0; i < stats.length; ++i) {
		if (stats[i] != 0 && (i < 5 || (i >= 35 && i <= 39)) && (i in indexToStatFunctionMap)) {
			tooltipList.push(indexToStatFunctionMap[i](stats[i]))
		}
	}
	if (tooltipList.length != 0) {
		return tooltipList.join('<br>')
	}
	return ""
}

function getBonusStatsTooltip(stats: any): string {
	const tooltipList: String[] = []
	for (let i = 0; i < stats.length; ++i) {
		if (stats[i] != 0 && (i > 5 && (i < 35 || i > 39)) && (i in indexToStatFunctionMap)) {
			if (i == 27 && stats[i] == stats[17]) {
				continue
			}
			tooltipList.push(indexToStatFunctionMap[i](stats[i]))
		}
	}
	if (tooltipList.length != 0) {
		return `<span style="color: #1eff00;">${tooltipList.join('<br>')}</span>`
	}
	return ""
}

function getWeaponSkillTooltip(item: any): string {
	const tooltipList: String[] = []
	for (let i = 0; i < item.weaponSkills.length; ++i) {
		if (item.weaponSkills[i] != 0 && (i in weaponSkillToStringFunctionMap)) {
			tooltipList.push(weaponSkillToStringFunctionMap[i](item.weaponSkills[i]))
		}
	}
	if (tooltipList.length != 0) {
		return `<span style="color: #1eff00;">${tooltipList.join('<br>')}</span>`
	}
	return ""
}

function getStrengthTooltip(val: number): string {
	return `+${val} Strength`
}

function getAgilityTooltip(val: number): string {
	return `+${val} Agility`
}

function getStaminaTooltip(val: number): string {
	return `+${val} Stamina`
}

function getIntellectTooltip(val: number): string {
	return `+${val} Intellect`
}

function getSpiritTooltip(val: number): string {
	return `+${val} Spirit`
}

function getSpellPowerTooltip(val: number): string {
	return `Equip: Increases damage and healing done by magical spells and effects by up to ${val}.`
}

function getArcanePowerTooltip(val: number): string {
	return `Equip: Increases damage done by Arcane spells and effects by up to ${val}.`
}

function getFirePowerTooltip(val: number): string {
	return `Equip: Increases damage done by Fire spells and effects by up to ${val}.`
}

function getFrostPowerTooltip(val: number): string {
	return `Equip: Increases damage done by Frost spells and effects by up to ${val}.`
}

function getHolyPowerTooltip(val: number): string {
	return `Equip: Increases damage done by Holy spells and effects by up to ${val}.`
}

function getNaturePowerTooltip(val: number): string {
	return `Equip: Increases damage done by Nature spells and effects by up to ${val}.`
}

function getShadowPowerTooltip(val: number): string {
	return `Equip: Increases damage done by Shadow spells and effects by up to ${val}.`
}

function getMP5Tooltip(val: number): string {
	return `Equip: Restores ${val} mana per 5 sec.`
}

function getSpellHitTooltip(val: number): string {
	return `Equip: Improves your chance to hit with spells by ${val}%.`
}

function getSpellCritTooltip(val: number): string {
	return `Equip: Improves your chance to get a critical strike with spells by ${val}%.`
}

function getSpellHasteTooltip(val: number): string {
	return `Equip: Increases your attack and casting speed by ${val}%.`
}

function getSpellPenetrationTooltip(val: number): string {
	return `Equip: Decreases the magical resistances of your spell targets by ${val}.`
}

function getAttackPowerTooltip(val: number): string {
	return `Equip: +${val} Attack Power.`
}

function getMeleeHitTooltip(val: number): string {
	return `Equip: Improves your chance to hit by ${val}%.`
}

function getMeleeCritTooltip(val: number): string {
	return `Equip: Improves your chance to get a critical strike by ${val}%.`
}

function getMeleeHasteTooltip(val: number): string {
	return `Equip: Increases your attack and casting speed by ${val}%.`
}

function getArmorPenetrationTooltip(val: number): string {
	return `Equip: Your attacks ignore ${val} of the target's armor.`
}

function getRangedAttackPowerTooltip(val: number): string {
	return `Equip: +${val} ranged Attack Power.`
}

function getDefenseTooltip(val: number): string {
	return `Equip: Increased Defense +${val}.`
}

function getBlockTooltip(val: number): string {
	return `Equip: Increases your chance to block attacks with a shield by ${val}%.`
}

function getBlockValueTooltip(val: number): string {
	return `Equip: Increases the block value of your shield by ${val}.`
}

function getDodgeTooltip(val: number): string {
	return `Equip: Increases your chance to dodge an attack by ${val}%.`
}

function getParryTooltip(val: number): string {
	return `Equip: Increases your chance to parry an attack by ${val}%.`
}

function getArcaneResistanceTooltip(val: number): string {
	return `+${val} Arcane Resistance`
}

function getFireResistanceTooltip(val: number): string {
	return `+${val} Fire Resistance`
}

function getFrostResistanceTooltip(val: number): string {
	return `+${val} Frost Resistance`
}

function getNatureResistanceTooltip(val: number): string {
	return `+${val} Nature Resistance`
}

function getShadowResistanceTooltip(val: number): string {
	return `+${val} Shadow Resistance`
}

function getHealingPowerTooltip(val: number): string {
	return `Equip: Increases healing done by spells and effects by up to ${val}.`
}

function getSpellDamageTooltip(val: number): string {
	return `Equip: Increases damage done by magical spells and effects by up to ${val}.`
}

function getFeralAttackPowerTooltip(val: number): string {
	return `Equip: +${val} Attack Power in Cat, Bear, Dire Bear, and Moonkin forms only.`
}

function getAxesWeaponSkillTooltip(val: number): string {
	return `Equip: Increased Axes +${val}.`
}

function getSwordsWeaponSkillTooltip(val: number): string {
	return `Equip: Increased Swords +${val}.`
}

function getMacesWeaponSkillTooltip(val: number): string {
	return `Equip: Increased Maces +${val}.`
}

function getDaggersWeaponSkillTooltip(val: number): string {
	return `Equip: Increased Daggers +${val}.`
}

function getUnarmedWeaponSkillTooltip(val: number): string {
	return `Equip: Increased Unarmed +${val}.`
}

function getTwoHandedAxesWeaponSkillTooltip(val: number): string {
	return `Equip: Increased Two-Handed Axes +${val}.`
}

function getTwoHandedSwordsWeaponSkillTooltip(val: number): string {
	return `Equip: Increased Two-Handed Swords +${val}.`
}

function getTwoHandedMacesWeaponSkillTooltip(val: number): string {
	return `Equip: Increased Two-Handed Maces +${val}.`
}

function getPolearmsWeaponSkillTooltip(val: number): string {
	return `Equip: Increased Polearms +${val}.`
}

function getStavesWeaponSkillTooltip(val: number): string {
	return `Equip: Increased Staves +${val}.`
}

function getThrownWeaponSkillTooltip(val: number): string {
	return `Equip: Increased Thrown +${val}.`
}

function getBowsWeaponSkillTooltip(val: number): string {
	return `Equip: Increased Bows +${val}.`
}

function getCrossbowsWeaponSkillTooltip(val: number): string {
	return `Equip: Increased Crossbows +${val}.`
}

function getGunsWeaponSkillTooltip(val: number): string {
	return `Equip: Increased Guns +${val}.`
}