package ies

import "rrc/utils"

// UL-PDCP-ExcessDelayResult-r17 ::= SEQUENCE
// Extensible
type UlPdcpExcessdelayresultR17 struct {
	DrbIdR17       DrbIdentity
	ExcessdelayR17 utils.INTEGER `lb:0,ub:31`
}
