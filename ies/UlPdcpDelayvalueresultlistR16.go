package ies

import "rrc/utils"

// UL-PDCP-DelayValueResultList-r16 ::= SEQUENCE OF UL-PDCP-DelayValueResult-r16
// SIZE (1..maxDRB)
type UlPdcpDelayvalueresultlistR16 struct {
	Value utils.Sequence[UlPdcpDelayvalueresultR16]
}
