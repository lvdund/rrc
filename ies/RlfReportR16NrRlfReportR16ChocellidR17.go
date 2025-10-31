package ies

// RLF-Report-r16-nr-RLF-Report-r16-choCellId-r17 ::= CHOICE
const (
	RlfReportR16NrRlfReportR16ChocellidR17ChoiceNothing = iota
	RlfReportR16NrRlfReportR16ChocellidR17ChoiceCellglobalidR17
	RlfReportR16NrRlfReportR16ChocellidR17ChoicePciArfcnR17
)

type RlfReportR16NrRlfReportR16ChocellidR17 struct {
	Choice          uint64
	CellglobalidR17 *CgiInfoLoggingR16
	PciArfcnR17     *PciArfcnNrR16
}
