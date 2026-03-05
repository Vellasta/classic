package core

import (
	"github.com/Vellasta/classic/sim/core/proto"
	"github.com/Vellasta/classic/sim/core/stats"
)

type APLValueCurrentAttackPower struct {
	DefaultAPLValueImpl
	unit *Unit
}

func (rot *APLRotation) newValueCurrentAttackPower(_ *proto.APLValueCurrentAttackPower) APLValue {
	return &APLValueCurrentAttackPower{
		unit: rot.unit,
	}
}
func (value *APLValueCurrentAttackPower) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeFloat
}
func (value *APLValueCurrentAttackPower) GetFloat(_ *Simulation) float64 {
	return value.unit.GetStat(stats.AttackPower)
}
func (value *APLValueCurrentAttackPower) String() string {
	return "Current Attack Power"
}
