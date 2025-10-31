package ies

// PosMeasGapPreConfigToAddModList-r17 ::= SEQUENCE OF PosGapConfig-r17
// SIZE (1..maxNrofPreConfigPosGapId-r17)
type PosmeasgappreconfigtoaddmodlistR17 struct {
	Value []PosgapconfigR17 `lb:1,ub:maxNrofPreConfigPosGapIdR17`
}
