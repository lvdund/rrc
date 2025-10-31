package ies

// VisitedCellInfo-r16-visitedCellId-r16-nr-CellId-r16 ::= CHOICE
const (
	VisitedcellinfoR16VisitedcellidR16NrCellidR16ChoiceNothing = iota
	VisitedcellinfoR16VisitedcellidR16NrCellidR16ChoiceCgiInfo
	VisitedcellinfoR16VisitedcellidR16NrCellidR16ChoicePciArfcnR16
)

type VisitedcellinfoR16VisitedcellidR16NrCellidR16 struct {
	Choice      uint64
	CgiInfo     *CgiInfoLoggingR16
	PciArfcnR16 *PciArfcnNrR16
}
