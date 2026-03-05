package guardians

import "github.com/Vellasta/classic/sim/core"

func ConstructGuardians(character *core.Character) {
	constructEmeralDragonWhelps(character)
	constructEskhandar(character)
	constructCoreHound(character)
}
