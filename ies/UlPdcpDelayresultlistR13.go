package ies

import "rrc/utils"

// UL-PDCP-DelayResultList-r13 ::= SEQUENCE OF UL-PDCP-DelayResult-r13
// SIZE (1..maxQCI-r13)
type UlPdcpDelayresultlistR13 struct {
	Value utils.Sequence[UlPdcpDelayresultR13]
}
