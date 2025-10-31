package ies

import "rrc/utils"

// BandNR-cancelOverlappingPUSCH-r16 ::= ENUMERATED
type BandnrCanceloverlappingpuschR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrCanceloverlappingpuschR16EnumeratedNothing = iota
	BandnrCanceloverlappingpuschR16EnumeratedSupported
)
