package sim

import (
	_ "github.com/Vellasta/classic/sim/common"
	"github.com/Vellasta/classic/sim/druid/balance"
	"github.com/Vellasta/classic/sim/paladin/retribution"
	dpsrogue "github.com/Vellasta/classic/sim/rogue/dps_rogue"
	"github.com/Vellasta/classic/sim/shaman/elemental"
	"github.com/Vellasta/classic/sim/shaman/enhancement"
	"github.com/Vellasta/classic/sim/shaman/warden"

	"github.com/Vellasta/classic/sim/druid/feral"
	// restoDruid "github.com/Vellasta/classic/sim/druid/restoration"
	// feralTank "github.com/Vellasta/classic/sim/druid/tank"
	_ "github.com/Vellasta/classic/sim/encounters"
	"github.com/Vellasta/classic/sim/hunter"
	"github.com/Vellasta/classic/sim/mage"

	// holyPaladin "github.com/Vellasta/classic/sim/paladin/holy"
	"github.com/Vellasta/classic/sim/paladin/protection"
	// "github.com/Vellasta/classic/sim/paladin/retribution"
	// healingPriest "github.com/Vellasta/classic/sim/priest/healing"
	"github.com/Vellasta/classic/sim/priest/shadow"

	// restoShaman "github.com/Vellasta/classic/sim/shaman/restoration"
	dpsWarlock "github.com/Vellasta/classic/sim/warlock/dps"
	dpsWarrior "github.com/Vellasta/classic/sim/warrior/dps_warrior"
	tankWarrior "github.com/Vellasta/classic/sim/warrior/tank_warrior"
)

var registered = false

func RegisterAll() {
	if registered {
		return
	}
	registered = true

	balance.RegisterBalanceDruid()
	feral.RegisterFeralDruid()
	// feralTank.RegisterFeralTankDruid()
	// restoDruid.RegisterRestorationDruid()
	elemental.RegisterElementalShaman()
	enhancement.RegisterEnhancementShaman()
	warden.RegisterWardenShaman()
	// restoShaman.RegisterRestorationShaman()
	hunter.RegisterHunter()
	mage.RegisterMage()
	// healingPriest.RegisterHealingPriest()
	shadow.RegisterShadowPriest()
	dpsrogue.RegisterDpsRogue()
	dpsWarrior.RegisterDpsWarrior()
	tankWarrior.RegisterTankWarrior()
	// holyPaladin.RegisterHolyPaladin()
	protection.RegisterProtectionPaladin()
	retribution.RegisterRetributionPaladin()
	dpsWarlock.RegisterDpsWarlock()
}
