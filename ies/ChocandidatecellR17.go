package ies

// ChoCandidateCell-r17 ::= CHOICE
const (
	ChocandidatecellR17ChoiceNothing = iota
	ChocandidatecellR17ChoiceCellglobalidR17
	ChocandidatecellR17ChoicePciArfcnR17
)

type ChocandidatecellR17 struct {
	Choice          uint64
	CellglobalidR17 *CgiInfoLoggingR16
	PciArfcnR17     *PciArfcnNrR16
}
