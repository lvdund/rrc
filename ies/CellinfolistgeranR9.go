package ies

import "rrc/utils"

// CellInfoListGERAN-r9 ::= SEQUENCE OF CellInfoGERAN-r9
// SIZE (1..maxCellInfoGERAN-r9)
type CellinfolistgeranR9 struct {
	Value []CellinfogeranR9
}
