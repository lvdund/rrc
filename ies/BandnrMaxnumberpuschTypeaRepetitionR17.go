package ies

import "rrc/utils"

// BandNR-maxNumberPUSCH-TypeA-Repetition-r17 ::= ENUMERATED
type BandnrMaxnumberpuschTypeaRepetitionR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrMaxnumberpuschTypeaRepetitionR17EnumeratedNothing = iota
	BandnrMaxnumberpuschTypeaRepetitionR17EnumeratedSupported
)
