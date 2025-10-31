package ies

// VisitedPSCellInfoList-r17 ::= SEQUENCE OF VisitedPSCellInfo-r17
// SIZE (1..maxPSCellHistory-r17)
type VisitedpscellinfolistR17 struct {
	Value []VisitedpscellinfoR17 `lb:1,ub:maxPSCellHistoryR17`
}
