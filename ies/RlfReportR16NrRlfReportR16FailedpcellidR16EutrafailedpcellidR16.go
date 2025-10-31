package ies

// RLF-Report-r16-nr-RLF-Report-r16-failedPCellId-r16-eutraFailedPCellId-r16 ::= CHOICE
const (
	RlfReportR16NrRlfReportR16FailedpcellidR16EutrafailedpcellidR16ChoiceNothing = iota
	RlfReportR16NrRlfReportR16FailedpcellidR16EutrafailedpcellidR16ChoiceCellglobalidR16
	RlfReportR16NrRlfReportR16FailedpcellidR16EutrafailedpcellidR16ChoicePciArfcnR16
)

type RlfReportR16NrRlfReportR16FailedpcellidR16EutrafailedpcellidR16 struct {
	Choice          uint64
	CellglobalidR16 *CgiInfoeutralogging
	PciArfcnR16     *PciArfcnEutraR16
}
