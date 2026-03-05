package encounters

import (
	"github.com/Vellasta/classic/sim/core"
)

func init() {
	// TODO: Classic encounters?
	// naxxramas.Register()
	addPatchwerk("Classic")
	addNothThePlaguebringer("Classic")
	addKelthuzad("Classic")
	addKeeperGnarlmoon("Classic")
	addKruul("Classic")
	addLevel60("Classic")
}

func AddSingleTargetBossEncounter(presetTarget *core.PresetTarget) {
	core.AddPresetTarget(presetTarget)
	core.AddPresetEncounter(presetTarget.Config.Name, []string{
		presetTarget.Path(),
	})
}
