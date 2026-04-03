import { ItemSlot } from '../../proto/common';
import { Player } from '../../player';
import { ItemListType } from './item_list';
import { Database } from '../../proto_utils/database';

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

const setIDToStringFunctionMap: Record<number, Function> = {
	[206]: getGiantstalkerTooltip,
	[215]: getDragonstalkerTooltip,
	[477]: getPredatorsTooltip,
	[515]: getBeastmasterTooltip,
	[510]: getTrappingsTooltip,
	[509]: getStrikersTooltip,
	[530]: getCryptstalkerTooltip,
	[697]: getRavenstalkerTooltip,
	[719]: getCombatantsTooltip,
	[720]: getPartisansTooltip,
	[721]: getVeteransTooltip,
}


export function createItemTooltip(item: ItemListType, player: Player<any>): string {
	const nameTooltip = getNameTooltip(item)
	const ilvlTooltip = getIlvlTooltip(item)
	const itemTypeTooltip = getitemTypeTooltip(item)
	const weaponTooltip = getWeaponTooltip(item)
	const armorTooltip = getArmorTooltip(item.stats)
	const baseStatsTooltip = getBaseStatsTooltip(item.stats)
	const bonusStatsTooltip = getBonusStatsTooltip(item.stats)
	const bonusWeaponDamageTooltip = getBonusWeaponDamageTooltip(item)
	const FortuneTooltip = getFortuneTooltip(item)
	const weaponSkillTooltip = getWeaponSkillTooltip(item)
	const procStringTooltip = getProcStringTooltip(item)
	const setTooltip = getSetTooltip(item, player)

	const tooltipList: string[] = [
		nameTooltip, 
		ilvlTooltip, 
		itemTypeTooltip,
		weaponTooltip, 
		armorTooltip, 
		baseStatsTooltip, 
		bonusStatsTooltip,
		bonusWeaponDamageTooltip,
		FortuneTooltip,
		weaponSkillTooltip,
		procStringTooltip,
		setTooltip
	]
	const filteredTooltipList = tooltipList.filter((s) => s != "")
	return filteredTooltipList.join('<br>')
}

function getNameTooltip(item: any): string {
	return `<span style="color: ${qualityToColorMap[item.quality]}; font-size: 112.5%;">${item.name}</span>`
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
		if (stats[i] != 0 && (i >= 5 && (i < 35 || i > 39)) && (i in indexToStatFunctionMap)) {
			if (i == 27 && stats[i] == stats[17]) {
				continue
			}
			if (i == 20 && stats[i] == stats[15]) {
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

function getBonusWeaponDamageTooltip(item: any): string {
	if (item.bonusPhysicalDamage > 0) {
		return `<span style="color: #1eff00;">Equip: +${item.bonusPhysicalDamage} Weapon Damage.</span>`
	}
	return ""
}

function getFortuneTooltip(item: any): string {
	if (item.fortune > 0) {
		return `<span style="color: #1eff00;">Equip: Increases your chance to trigger effects from equipped items by ${item.fortune}%.</span>`
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

function getProcStringTooltip(item: any): string {
	const tooltipList: String[] = []
	for (let i = 0; i < item.procStrings.length; ++i) {
		if (item.procStrings[i] != "") {
			tooltipList.push(item.procStrings[i])
		}
	}
	if (tooltipList.length != 0) {
		return `<span style="color: #1eff00;">${tooltipList.join('<br>')}</span>`
	}
	return ""
}

function getSetTooltip(item: any, player: Player<any>): string {
	if (item.setId in setIDToStringFunctionMap) {
		return `<br>${setIDToStringFunctionMap[item.setId](player)}`
	}
	return ""
}

function getSetCountAndPieces( 
	idToStringMap: Record<number, String>, 
	player: Player<any>, 
	setCountToStringMap: Record<number, String>
): {
	setCount: number, 
	setPieces: String[], 
	setString: String[]
} {
	const setPieces: String[] = []
	var playerEquipped = new Set<number>()
	for (let i = 0; i < 17; ++i) {
		const equipItem = player.getEquippedItem(i)
		if (equipItem) {
			playerEquipped.add(equipItem.item.id)
		}
	}
	
	var setCount = 0
	for (const [itemId, itemName] of Object.entries(idToStringMap)) {
		var setItemName
		if (playerEquipped.has(+itemId)) {
			setItemName = `<span style="color: #ffffad; padding-left: 10px;">${itemName}</span>`
			setCount += 1
		} else {
			setItemName = `<span style="color: #a7a7a7; padding-left: 10px;">${itemName}</span>`
		}
		setPieces.push(setItemName)
	}

	const setString: String[] = []
	for (const [setIndex, setValue] of Object.entries(setCountToStringMap)) {
		if (+setIndex <= setCount) {
			setString.push(`<span style="color: #1eff00;">${setValue}</span>`)
		} else {
			setString.push(`<span style="color: #a7a7a7;">${setValue}</span>`)
		}
	}

	return {setCount, setPieces, setString}
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

function getGiantstalkerTooltip(player: Player<any>): string {
	const idToStringMap: Record<number, String> = {
		[16851]: "Giantstalker's Belt",
		[16849]: "Giantstalker's Boots",
		[16850]: "Giantstalker's Bracers",
		[16845]: "Giantstalker's Breastplate",
		[16848]: "Giantstalker's Epaulets",
		[16852]: "Giantstalker's Gloves",
		[16846]: "Giantstalker's Helmet",
		[16847]: "Giantstalker's Leggings",
	}

	const setCountToStringMap: Record<number, String> = {
		[3]: "(3) Set : Increases the range of your Mend Pet spell by 50% and the effect by 10%.  Also reduces the cost by 30%.",
		[5]: "(5) Set : Increases your pet's stamina by 30 and all spell resistances by 40.",
		[8]: "(8) Set : Increases the damage of Multi-shot and Volley by 15%.",
	}

	const {setCount, setPieces, setString} = getSetCountAndPieces(idToStringMap, player, setCountToStringMap)
	
	return `<span style="color: gold;">Giantstalker Armor (${setCount}/8)</span><br>${setPieces.join('<br>')}<br><br>${setString.join('<br>')}`
}

function getDragonstalkerTooltip(player: Player<any>): string {
	const idToStringMap: Record<number, String> = {
		[16936]: "Dragonstalker's Belt",
		[16935]: "Dragonstalker's Bracers",
		[16942]: "Dragonstalker's Breastplate",
		[16940]: "Dragonstalker's Gauntlets",
		[16941]: "Dragonstalker's Greaves",
		[16939]: "Dragonstalker's Helm",
		[16938]: "Dragonstalker's Legguards",
		[16937]: "Dragonstalker's Spaulders",
	}

	const setCountToStringMap: Record<number, String> = {
		[3]: "(3) Set: Increases the Ranged Attack Power bonus of your Aspect of the Hawk by 20% and Melee Attack Power bonus of your Aspect of the Wolf by 20%.",
		[5]: "(5) Set : Increases your pet's stamina by 40 and all spell resistances by 60.",
		[8]: "(8) Set: You have a chance whenever you deal melee or ranged damage to gain the Detect Weakness effect. Detect Weakness increases your Attack Power by 450 for 7 sec.",
	}

	const {setCount, setPieces, setString} = getSetCountAndPieces(idToStringMap, player, setCountToStringMap)
	
	return `<span style="color: gold;">Dragonstalker Armor (${setCount}/8)</span><br>${setPieces.join('<br>')}<br><br>${setString.join('<br>')}`
}

function getPredatorsTooltip(player: Player<any>): string {
	const idToStringMap: Record<number, String> = {
		[19621]: "Maelstrom's Wrath",
		[19953]: "Renataki's Charm of Beasts",
		[19833]: "Zandalar Predator's Bracers",
		[19832]: "Zandalar Predator's Belt",
		[19831]: "Zandalar Predator's Mantle",
	}

	const setCountToStringMap: Record<number, String> = {
		[2]: "(2) Set : +20 Attack Power.",
		[3]: "(3) Set : Decreases the cooldown of Concussive Shot by 1 sec.",
		[5]: "(5) Set : Increases the duration of Serpent Sting by 3 sec.",
	}

	const {setCount, setPieces, setString} = getSetCountAndPieces(idToStringMap, player, setCountToStringMap)
	
	return `<span style="color: gold;">Predator's Armor (${setCount}/5)</span><br>${setPieces.join('<br>')}<br><br>${setString.join('<br>')}`
}

function getBeastmasterTooltip(player: Player<any>): string {
	const idToStringMap: Record<number, String> = {
		[22010]: "Beastmaster's Belt",
		[22011]: "Beastmaster's Bindings",
		[22061]: "Beastmaster's Boots",
		[22013]: "Beastmaster's Cap",
		[22015]: "Beastmaster's Gloves",
		[22016]: "Beastmaster's Mantle",
		[22017]: "Beastmaster's Pants",
		[22060]: "Beastmaster's Tunic",
	}

	const setCountToStringMap: Record<number, String> = {
		[2]: "(2) Set : +8 All Resistances.",
		[4]: "(4) Set : Your normal ranged attacks have a 4% chance of restoring 200 mana. (Proc chance: 4%)",
		[6]: "(6) Set : +40 Attack Power.",
		[8]: "(8) Set : +200 Armor.",
	}

	const {setCount, setPieces, setString} = getSetCountAndPieces(idToStringMap, player, setCountToStringMap)
	
	return `<span style="color: gold;">Beastmaster Armor (${setCount}/8)</span><br>${setPieces.join('<br>')}<br><br>${setString.join('<br>')}`
}

function getTrappingsTooltip(player: Player<any>): string {
	const idToStringMap: Record<number, String> = {
		[21401]: "Scythe of the Unseen Path",
		[21402]: "Signet of the Unseen Path",
		[21403]: "Cloak of the Unseen Path",
	}

	const setCountToStringMap: Record<number, String> = {
		[2]: "(2) Set: Increases your pet's damage by 3%.",
		[3]: "(3) Set: Increases pet Focus regeneration by 2.<br>(3) Set: Increases the critical strike chance of Steady Shot, Raptor Strike and Mongoose Bite by 2%.",
	}

	const {setCount, setPieces, setString} = getSetCountAndPieces(idToStringMap, player, setCountToStringMap)
	
	return `<span style="color: gold;">Trappings of the Unseen Path (${setCount}/3)</span><br>${setPieces.join('<br>')}<br><br>${setString.join('<br>')}`
}


function getStrikersTooltip(player: Player<any>): string {
	const idToStringMap: Record<number, String> = {
		[21366]: "Striker's Diadem",
		[21365]: "Striker's Footguards",
		[21370]: "Striker's Hauberk",
		[21368]: "Striker's Leggings",
		[21367]: "Striker's Pauldrons",
	}

	const setCountToStringMap: Record<number, String> = {
		[3]: "(3) Set : Reduces the cost of your Arcane Shots by 10%. (Proc chance: 20%)",
		[5]: "(5) Set : Reduces the cooldown of your Rapid Fire ability by 2 minutes. (Proc chance: 20%)",
	}

	const {setCount, setPieces, setString} = getSetCountAndPieces(idToStringMap, player, setCountToStringMap)
	
	return `<span style="color: gold;">Striker's Garb (${setCount}/5)</span><br>${setPieces.join('<br>')}<br><br>${setString.join('<br>')}`
}

function getCryptstalkerTooltip(player: Player<any>): string {
	const idToStringMap: Record<number, String> = {
		[22440]: "Cryptstalker Boots",
		[22442]: "Cryptstalker Girdle",
		[22441]: "Cryptstalker Handguards",
		[22438]: "Cryptstalker Headpiece",
		[22437]: "Cryptstalker Legguards",
		[22439]: "Cryptstalker Spaulders",
		[22436]: "Cryptstalker Tunic",
		[22443]: "Cryptstalker Wristguards",
		[23067]: "Ring of the Cryptstalker",
	}

	const setCountToStringMap: Record<number, String> = {
		[2]: "(2) Set : Increases the duration of your Rapid Fire by 4 secs.",
		[4]: "(4) Set : Increases Attack Power by 50 for both you and your pet.",
		[6]: "(6) Set : Your ranged critical hits cause an Adrenaline Rush, granting you 50 mana.",
		[8]: "(8) Set : Reduces the mana cost of your Multi-Shot and Aimed Shot by 20.",
	}

	const {setCount, setPieces, setString} = getSetCountAndPieces(idToStringMap, player, setCountToStringMap)
	
	return `<span style="color: gold;">Cryptstalker Armor (${setCount}/9)</span><br>${setPieces.join('<br>')}<br><br>${setString.join('<br>')}`
}

function getRavenstalkerTooltip(player: Player<any>): string {
	const idToStringMap: Record<number, String> = {
		[47318]: "Ravenstalker Headpiece",
		[47319]: "Ravenstalker Spaulders",
		[47320]: "Ravenstalker Tunic",
		[47321]: "Ravenstalker Legguards",
		[47322]: "Ravenstalker Boots",
		[47323]: "Ravenstalker Choker",
	}

	const setCountToStringMap: Record<number, String> = {
		[3]: "(3) Set: Reduces the cooldown of Multi-Shot and Carve by 1 sec.",
		[5]: "(5) Set: Increases the attack speed provided by Swift Aspects by an additional 5%.",
	}

	const {setCount, setPieces, setString} = getSetCountAndPieces(idToStringMap, player, setCountToStringMap)
	
	return `<span style="color: gold;">Ravenstalker Armor (${setCount}/6)</span><br>${setPieces.join('<br>')}<br><br>${setString.join('<br>')}`
}

function getCombatantsTooltip(player: Player<any>): string {
	const idToStringMap: Record<number, String> = {
		[33440]: "Combatant's Chain Helm",
		[33441]: "Combatant's Chain Spaulders",
		[33442]: "Combatant's Chain Breastplate",
		[33443]: "Combatant's Chain Grips",
		[33444]: "Combatant's Chain Legguards",
		[33445]: "Combatant's Chain Boots",
	}

	const setCountToStringMap: Record<number, String> = {
		[2]: "(2) Set: Increases the damage done by your Multi-Shot and Carve by 4%.",
		[4]: "(4) Set: Reduces the cooldown of Concussive Shot by 1000 sec and Wing Clip by 500 sec.",
		[6]: "(6) Set: +20 Stamina.<br>(6) Set: Reduces damage taken from critical hits and damage over time effects by 3%.",
	}

	const {setCount, setPieces, setString} = getSetCountAndPieces(idToStringMap, player, setCountToStringMap)
	
	return `<span style="color: gold;">Combatant's Pursuit (${setCount}/6)</span><br>${setPieces.join('<br>')}<br><br>${setString.join('<br>')}`
}

function getPartisansTooltip(player: Player<any>): string {
	const idToStringMap: Record<number, String> = {
		[33446]: "Partisan's Chain Helm",
		[33447]: "Partisan's Chain Spaulders",
		[33448]: "Partisan's Chain Breastplate",
		[33449]: "Partisan's Chain Grips",
		[33450]: "Partisan's Chain Legguards",
		[33451]: "Partisan's Chain Boots",
	}

	const setCountToStringMap: Record<number, String> = {
		[2]: "(2) Set: Increases the damage done by your Multi-Shot and Carve by 4%.",
		[4]: "(4) Set: Reduces the cooldown of Concussive Shot by 1000 sec and Wing Clip by 500 sec.",
		[6]: "(6) Set: +30 Stamina.<br>(6) Set: Reduces damage taken from critical hits and damage over time effects by 3%.",
	}

	const {setCount, setPieces, setString} = getSetCountAndPieces(idToStringMap, player, setCountToStringMap)
	
	return `<span style="color: gold;">Partisan's Pursuit (${setCount}/6)</span><br>${setPieces.join('<br>')}<br><br>${setString.join('<br>')}`
}

function getVeteransTooltip(player: Player<any>): string {
	const idToStringMap: Record<number, String> = {
		[33452]: "Veteran's Chain Helm",
		[33453]: "Veteran's Chain Spaulders",
		[33454]: "Veteran's Chain Breastplate",
		[33455]: "Veteran's Chain Grips",
		[33456]: "Veteran's Chain Legguards",
		[33457]: "Veteran's Chain Boots",
	}

	const setCountToStringMap: Record<number, String> = {
		[2]: "(2) Set: Increases the damage done by your Multi-Shot and Carve by 4%.",
		[4]: "(4) Set: Reduces the cooldown of Concussive Shot by 1000 sec and Wing Clip by 500 sec.",
		[6]: "(6) Set: +35 Stamina.<br>(6) Set: Reduces damage taken from critical hits and damage over time effects by 6%.",
	}

	const {setCount, setPieces, setString} = getSetCountAndPieces(idToStringMap, player, setCountToStringMap)
	
	return `<span style="color: gold;">Veteran's Pursuit (${setCount}/6)</span><br>${setPieces.join('<br>')}<br><br>${setString.join('<br>')}`
}

export async function getTooltipString(itemId: number, spellId: number): Promise<string> {
	// const url = `https://database.turtlecraft.gg/ajax.php?spell=${id}`;
	const url = itemId !== 0 ? `https://nether.wowhead.com/classic/tooltip/item/${itemId}` : `https://nether.wowhead.com/classic/tooltip/spell/${spellId}`
	try {
		const response = await fetch(url);
		const json = await response.json();
		if (!json['tooltip']) {
			const itemIcon = await Database.getItemIconData(itemId);
			if (itemIcon) {
				return itemIcon.tooltip
			}
			return "Item not found"
		}
		return json['tooltip'];
	} catch (e) {
		console.error('Error while fetching url: ' + url + '\n\n' + e);
		return "Error while fetching url";
	}
}