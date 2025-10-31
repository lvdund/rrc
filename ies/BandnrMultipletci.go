package ies

import "rrc/utils"

// BandNR-multipleTCI ::= ENUMERATED
type BandnrMultipletci struct {
	Value utils.ENUMERATED
}

const (
	BandnrMultipletciEnumeratedNothing = iota
	BandnrMultipletciEnumeratedSupported
)
