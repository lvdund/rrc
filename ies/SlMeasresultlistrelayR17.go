package ies

// SL-MeasResultListRelay-r17 ::= SEQUENCE OF SL-MeasResultRelay-r17
// SIZE (1..maxNrofRelayMeas-r17)
type SlMeasresultlistrelayR17 struct {
	Value []SlMeasresultrelayR17 `lb:1,ub:maxNrofRelayMeasR17`
}
