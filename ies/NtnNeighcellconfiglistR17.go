package ies

// NTN-NeighCellConfigList-r17 ::= SEQUENCE OF NTN-NeighCellConfig-r17
// SIZE (1..maxCellNTN-r17)
type NtnNeighcellconfiglistR17 struct {
	Value []NtnNeighcellconfigR17 `lb:1,ub:maxCellNTNR17`
}
