package ies

import "rrc/utils"

// VisitedCellInfoList-r12 ::= SEQUENCE OF VisitedCellInfo-r12
// SIZE (1..maxCellHistory-r12)
type VisitedcellinfolistR12 struct {
	Value []VisitedcellinfoR12
}
