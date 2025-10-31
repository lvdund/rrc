package ies

import "rrc/utils"

// BandCombination-v1560-ne-DC-BC ::= ENUMERATED
type BandcombinationV1560NeDcBc struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationV1560NeDcBcEnumeratedNothing = iota
	BandcombinationV1560NeDcBcEnumeratedSupported
)
