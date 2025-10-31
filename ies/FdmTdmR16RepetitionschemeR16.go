package ies

import "rrc/utils"

// FDM-TDM-r16-repetitionScheme-r16 ::= ENUMERATED
type FdmTdmR16RepetitionschemeR16 struct {
	Value utils.ENUMERATED
}

const (
	FdmTdmR16RepetitionschemeR16EnumeratedNothing = iota
	FdmTdmR16RepetitionschemeR16EnumeratedFdmschemea
	FdmTdmR16RepetitionschemeR16EnumeratedFdmschemeb
	FdmTdmR16RepetitionschemeR16EnumeratedTdmschemea
)
