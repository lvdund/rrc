package ies

import "rrc/utils"

// CellGroupConfig-uplinkTxSwitching-DualUL-TxState-r17 ::= ENUMERATED
type CellgroupconfigUplinktxswitchingDualulTxstateR17 struct {
	Value utils.ENUMERATED
}

const (
	CellgroupconfigUplinktxswitchingDualulTxstateR17EnumeratedNothing = iota
	CellgroupconfigUplinktxswitchingDualulTxstateR17EnumeratedOnet
	CellgroupconfigUplinktxswitchingDualulTxstateR17EnumeratedTwot
)
