package ies

// RLF-Report-r16 ::= CHOICE
// Extensible
const (
	RlfReportR16ChoiceNothing = iota
	RlfReportR16ChoiceNrRlfReportR16
	RlfReportR16ChoiceEutraRlfReportR16
)

type RlfReportR16 struct {
	Choice            uint64
	NrRlfReportR16    *RlfReportR16NrRlfReportR16
	EutraRlfReportR16 *RlfReportR16EutraRlfReportR16
}
