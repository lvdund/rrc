package ies

// CellInfoListGERAN-r9 ::= SEQUENCE OF CellInfoGERAN-r9
// SIZE (1..maxCellInfoGERAN-r9)
type CellinfolistgeranR9 struct {
	Value []CellinfogeranR9 `lb:1,ub:maxCellInfoGERANR9`
}
