package ies

import "rrc/utils"

// PosSchedulingInfo-r16 ::= SEQUENCE
// Extensible
type PosschedulinginfoR16 struct {
	OffsettosiUsedR16       *utils.BOOLEAN
	PossiPeriodicityR16     PosschedulinginfoR16PossiPeriodicityR16
	PossiBroadcaststatusR16 PosschedulinginfoR16PossiBroadcaststatusR16
	PossibMappinginfoR16    PossibMappinginfoR16
}
