package ies

// VisitedPSCellInfo-r17-visitedCellId-r17-eutra-CellId-r17 ::= CHOICE
const (
	VisitedpscellinfoR17VisitedcellidR17EutraCellidR17ChoiceNothing = iota
	VisitedpscellinfoR17VisitedcellidR17EutraCellidR17ChoiceCellglobalidR17
	VisitedpscellinfoR17VisitedcellidR17EutraCellidR17ChoicePciArfcnR17
)

type VisitedpscellinfoR17VisitedcellidR17EutraCellidR17 struct {
	Choice          uint64
	CellglobalidR17 *CgiInfoeutralogging
	PciArfcnR17     *PciArfcnEutraR16
}
