package ies

// RLF-Report-r16-nr-RLF-Report-r16-failedPCellId-r16-nrFailedPCellId-r16 ::= CHOICE
const (
	RlfReportR16NrRlfReportR16FailedpcellidR16NrfailedpcellidR16ChoiceNothing = iota
	RlfReportR16NrRlfReportR16FailedpcellidR16NrfailedpcellidR16ChoiceCellglobalidR16
	RlfReportR16NrRlfReportR16FailedpcellidR16NrfailedpcellidR16ChoicePciArfcnR16
)

type RlfReportR16NrRlfReportR16FailedpcellidR16NrfailedpcellidR16 struct {
	Choice          uint64
	CellglobalidR16 *CgiInfoLoggingR16
	PciArfcnR16     *PciArfcnNrR16
}
