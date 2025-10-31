package ies

import "rrc/utils"

// BandNR-bwp-WithoutRestriction ::= ENUMERATED
type BandnrBwpWithoutrestriction struct {
	Value utils.ENUMERATED
}

const (
	BandnrBwpWithoutrestrictionEnumeratedNothing = iota
	BandnrBwpWithoutrestrictionEnumeratedSupported
)
