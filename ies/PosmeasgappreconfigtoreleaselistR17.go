package ies

// PosMeasGapPreConfigToReleaseList-r17 ::= SEQUENCE OF MeasPosPreConfigGapId-r17
// SIZE (1..maxNrofPreConfigPosGapId-r17)
type PosmeasgappreconfigtoreleaselistR17 struct {
	Value []MeaspospreconfiggapidR17 `lb:1,ub:maxNrofPreConfigPosGapIdR17`
}
