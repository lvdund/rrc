package ies

import "rrc/utils"

// RACH-Skip-r14-ul-ConfigInfo-r14-ul-SchedInterval-r14 ::= ENUMERATED
type RachSkipR14UlConfiginfoR14UlSchedintervalR14 struct {
	Value utils.ENUMERATED
}

const (
	RachSkipR14UlConfiginfoR14UlSchedintervalR14EnumeratedNothing = iota
	RachSkipR14UlConfiginfoR14UlSchedintervalR14EnumeratedSf2
	RachSkipR14UlConfiginfoR14UlSchedintervalR14EnumeratedSf5
	RachSkipR14UlConfiginfoR14UlSchedintervalR14EnumeratedSf10
)
