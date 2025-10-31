package ies

// VisitedPSCellInfo-r17-visitedCellId-r17-nr-CellId-r17 ::= CHOICE
const (
	VisitedpscellinfoR17VisitedcellidR17NrCellidR17ChoiceNothing = iota
	VisitedpscellinfoR17VisitedcellidR17NrCellidR17ChoiceCgiInfoR17
	VisitedpscellinfoR17VisitedcellidR17NrCellidR17ChoicePciArfcnR17
)

type VisitedpscellinfoR17VisitedcellidR17NrCellidR17 struct {
	Choice      uint64
	CgiInfoR17  *CgiInfoLoggingR16
	PciArfcnR17 *PciArfcnNrR16
}
