package ies

// RLF-Report-r16-nr-RLF-Report-r16-reconnectCellId-r16 ::= CHOICE
const (
	RlfReportR16NrRlfReportR16ReconnectcellidR16ChoiceNothing = iota
	RlfReportR16NrRlfReportR16ReconnectcellidR16ChoiceNrreconnectcellidR16
	RlfReportR16NrRlfReportR16ReconnectcellidR16ChoiceEutrareconnectcellidR16
)

type RlfReportR16NrRlfReportR16ReconnectcellidR16 struct {
	Choice                  uint64
	NrreconnectcellidR16    *CgiInfoLoggingR16
	EutrareconnectcellidR16 *CgiInfoeutralogging
}
