package ies

// UL-PDCP-DelayResultList-r13 ::= SEQUENCE OF UL-PDCP-DelayResult-r13
// SIZE (1..maxQCI-r13)
type UlPdcpDelayresultlistR13 struct {
	Value []UlPdcpDelayresultR13 `lb:1,ub:maxQCIR13`
}
