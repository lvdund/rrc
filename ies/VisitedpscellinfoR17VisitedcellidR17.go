package ies

// VisitedPSCellInfo-r17-visitedCellId-r17 ::= CHOICE
const (
	VisitedpscellinfoR17VisitedcellidR17ChoiceNothing = iota
	VisitedpscellinfoR17VisitedcellidR17ChoiceNrCellidR17
	VisitedpscellinfoR17VisitedcellidR17ChoiceEutraCellidR17
)

type VisitedpscellinfoR17VisitedcellidR17 struct {
	Choice         uint64
	NrCellidR17    *VisitedpscellinfoR17VisitedcellidR17NrCellidR17
	EutraCellidR17 *VisitedpscellinfoR17VisitedcellidR17EutraCellidR17
}
