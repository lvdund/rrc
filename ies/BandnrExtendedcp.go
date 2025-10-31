package ies

import "rrc/utils"

// BandNR-extendedCP ::= ENUMERATED
type BandnrExtendedcp struct {
	Value utils.ENUMERATED
}

const (
	BandnrExtendedcpEnumeratedNothing = iota
	BandnrExtendedcpEnumeratedSupported
)
