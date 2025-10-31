package ies

import "rrc/utils"

// UL-PDCP-DelayResult-r13 ::= SEQUENCE
// Extensible
type UlPdcpDelayresultR13 struct {
	QciIdR13       UlPdcpDelayresultR13QciIdR13
	ExcessdelayR13 utils.INTEGER `lb:0,ub:31`
}
