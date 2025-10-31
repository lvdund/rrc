package ies

import "rrc/utils"

// BandNR-cg-SDT-r17 ::= ENUMERATED
type BandnrCgSdtR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrCgSdtR17EnumeratedNothing = iota
	BandnrCgSdtR17EnumeratedSupported
)
