package ies

import "rrc/utils"

// CellGroupConfig-uplinkTxSwitchingOption-r16 ::= ENUMERATED
type CellgroupconfigUplinktxswitchingoptionR16 struct {
	Value utils.ENUMERATED
}

const (
	CellgroupconfigUplinktxswitchingoptionR16EnumeratedNothing = iota
	CellgroupconfigUplinktxswitchingoptionR16EnumeratedSwitchedul
	CellgroupconfigUplinktxswitchingoptionR16EnumeratedDualul
)
