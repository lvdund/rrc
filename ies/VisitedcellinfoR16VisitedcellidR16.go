package ies

// VisitedCellInfo-r16-visitedCellId-r16 ::= CHOICE
const (
	VisitedcellinfoR16VisitedcellidR16ChoiceNothing = iota
	VisitedcellinfoR16VisitedcellidR16ChoiceNrCellidR16
	VisitedcellinfoR16VisitedcellidR16ChoiceEutraCellidR16
)

type VisitedcellinfoR16VisitedcellidR16 struct {
	Choice         uint64
	NrCellidR16    *VisitedcellinfoR16VisitedcellidR16NrCellidR16
	EutraCellidR16 *VisitedcellinfoR16VisitedcellidR16EutraCellidR16
}
