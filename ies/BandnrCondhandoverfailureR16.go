package ies

import "rrc/utils"

// BandNR-condHandoverFailure-r16 ::= ENUMERATED
type BandnrCondhandoverfailureR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrCondhandoverfailureR16EnumeratedNothing = iota
	BandnrCondhandoverfailureR16EnumeratedSupported
)
