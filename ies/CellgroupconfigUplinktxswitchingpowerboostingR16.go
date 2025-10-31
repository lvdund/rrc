package ies

import "rrc/utils"

// CellGroupConfig-uplinkTxSwitchingPowerBoosting-r16 ::= ENUMERATED
type CellgroupconfigUplinktxswitchingpowerboostingR16 struct {
	Value utils.ENUMERATED
}

const (
	CellgroupconfigUplinktxswitchingpowerboostingR16EnumeratedNothing = iota
	CellgroupconfigUplinktxswitchingpowerboostingR16EnumeratedEnabled
)
