package ies

// UL-PDCP-DelayValueResultList-r16 ::= SEQUENCE OF UL-PDCP-DelayValueResult-r16
// SIZE (1..maxDRB)
type UlPdcpDelayvalueresultlistR16 struct {
	Value []UlPdcpDelayvalueresultR16 `lb:1,ub:maxDRB`
}
