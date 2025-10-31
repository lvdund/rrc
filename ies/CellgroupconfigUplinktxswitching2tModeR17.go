package ies

import "rrc/utils"

// CellGroupConfig-uplinkTxSwitching-2T-Mode-r17 ::= ENUMERATED
type CellgroupconfigUplinktxswitching2tModeR17 struct {
	Value utils.ENUMERATED
}

const (
	CellgroupconfigUplinktxswitching2tModeR17EnumeratedNothing = iota
	CellgroupconfigUplinktxswitching2tModeR17EnumeratedEnabled
)
