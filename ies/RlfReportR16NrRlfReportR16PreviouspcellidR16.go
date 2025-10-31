package ies

// RLF-Report-r16-nr-RLF-Report-r16-previousPCellId-r16 ::= CHOICE
const (
	RlfReportR16NrRlfReportR16PreviouspcellidR16ChoiceNothing = iota
	RlfReportR16NrRlfReportR16PreviouspcellidR16ChoiceNrpreviouscellR16
	RlfReportR16NrRlfReportR16PreviouspcellidR16ChoiceEutrapreviouscellR16
)

type RlfReportR16NrRlfReportR16PreviouspcellidR16 struct {
	Choice               uint64
	NrpreviouscellR16    *CgiInfoLoggingR16
	EutrapreviouscellR16 *CgiInfoeutralogging
}
