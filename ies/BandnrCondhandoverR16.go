package ies

import "rrc/utils"

// BandNR-condHandover-r16 ::= ENUMERATED
type BandnrCondhandoverR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrCondhandoverR16EnumeratedNothing = iota
	BandnrCondhandoverR16EnumeratedSupported
)
