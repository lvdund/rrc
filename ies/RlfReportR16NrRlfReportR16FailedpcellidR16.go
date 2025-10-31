package ies

// RLF-Report-r16-nr-RLF-Report-r16-failedPCellId-r16 ::= CHOICE
const (
	RlfReportR16NrRlfReportR16FailedpcellidR16ChoiceNothing = iota
	RlfReportR16NrRlfReportR16FailedpcellidR16ChoiceNrfailedpcellidR16
	RlfReportR16NrRlfReportR16FailedpcellidR16ChoiceEutrafailedpcellidR16
)

type RlfReportR16NrRlfReportR16FailedpcellidR16 struct {
	Choice                uint64
	NrfailedpcellidR16    *RlfReportR16NrRlfReportR16FailedpcellidR16NrfailedpcellidR16
	EutrafailedpcellidR16 *RlfReportR16NrRlfReportR16FailedpcellidR16EutrafailedpcellidR16
}
