package ies

// VisitedCellInfoList-r16 ::= SEQUENCE OF VisitedCellInfo-r16
// SIZE (1..maxCellHistory-r16)
type VisitedcellinfolistR16 struct {
	Value []VisitedcellinfoR16 `lb:1,ub:maxCellHistoryR16`
}
