package ies

import "rrc/utils"

// RACH-Skip-r14-ul-ConfigInfo-r14 ::= SEQUENCE
type RachSkipR14UlConfiginfoR14 struct {
	NumberofconfulProcessesR14 utils.INTEGER `lb:0,ub:8`
	UlSchedintervalR14         RachSkipR14UlConfiginfoR14UlSchedintervalR14
	UlStartsubframeR14         utils.INTEGER   `lb:0,ub:9`
	UlGrantR14                 utils.BITSTRING `lb:16,ub:16`
}
