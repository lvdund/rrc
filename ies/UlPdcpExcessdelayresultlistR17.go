package ies

// UL-PDCP-ExcessDelayResultList-r17 ::= SEQUENCE OF UL-PDCP-ExcessDelayResult-r17
// SIZE (1..maxDRB)
type UlPdcpExcessdelayresultlistR17 struct {
	Value []UlPdcpExcessdelayresultR17 `lb:1,ub:maxDRB`
}
