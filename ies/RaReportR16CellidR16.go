package ies

// RA-Report-r16-cellId-r16 ::= CHOICE
const (
	RaReportR16CellidR16ChoiceNothing = iota
	RaReportR16CellidR16ChoiceCellglobalidR16
	RaReportR16CellidR16ChoicePciArfcnR16
)

type RaReportR16CellidR16 struct {
	Choice          uint64
	CellglobalidR16 *CgiInfoLoggingR16
	PciArfcnR16     *PciArfcnNrR16
}
