package ies

import "rrc/utils"

// DummyD-amplitudeSubsetRestriction ::= ENUMERATED
type DummydAmplitudesubsetrestriction struct {
	Value utils.ENUMERATED
}

const (
	DummydAmplitudesubsetrestrictionEnumeratedNothing = iota
	DummydAmplitudesubsetrestrictionEnumeratedSupported
)
