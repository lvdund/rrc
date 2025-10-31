package ies

// VisitedCellInfo-r16-visitedCellId-r16-eutra-CellId-r16 ::= CHOICE
const (
	VisitedcellinfoR16VisitedcellidR16EutraCellidR16ChoiceNothing = iota
	VisitedcellinfoR16VisitedcellidR16EutraCellidR16ChoiceCellglobalidR16
	VisitedcellinfoR16VisitedcellidR16EutraCellidR16ChoicePciArfcnR16
)

type VisitedcellinfoR16VisitedcellidR16EutraCellidR16 struct {
	Choice          uint64
	CellglobalidR16 *CgiInfoeutra
	PciArfcnR16     *PciArfcnEutraR16
}
