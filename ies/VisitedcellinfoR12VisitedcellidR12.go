package ies

// VisitedCellInfo-r12-visitedCellId-r12 ::= CHOICE
const (
	VisitedcellinfoR12VisitedcellidR12ChoiceNothing = iota
	VisitedcellinfoR12VisitedcellidR12ChoiceCellglobalidR12
	VisitedcellinfoR12VisitedcellidR12ChoicePciArfcnR12
)

type VisitedcellinfoR12VisitedcellidR12 struct {
	Choice          uint64
	CellglobalidR12 *Cellglobalideutra
	PciArfcnR12     *VisitedcellinfoR12VisitedcellidR12PciArfcnR12
}
