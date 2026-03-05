package paladin

import (
	"time"

	"github.com/Vellasta/classic/sim/core"
	"github.com/Vellasta/classic/sim/core/stats"
)

// Libram IDs
const (
	SanctifiedOrb  = 20512
	LibramOfHope   = 22401
	LibramOfFervor = 23203
)

func init() {
	core.NewSimpleStatOffensiveTrinketEffect(SanctifiedOrb, stats.Stats{stats.MeleeCrit: 3 * core.CritRatingPerCritChance, stats.SpellCrit: 3 * core.CritRatingPerCritChance}, time.Second*25, time.Minute*3)
}
