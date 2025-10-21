package ies

import "rrc/utils"

// MeasResultListEUTRA ::= SEQUENCE OF MeasResultEUTRA
// SIZE (1..maxCellReport)
type Measresultlisteutra struct {
	Value []Measresulteutra
}
