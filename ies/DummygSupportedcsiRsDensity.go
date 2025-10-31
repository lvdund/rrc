package ies

import "rrc/utils"

// DummyG-supportedCSI-RS-Density ::= ENUMERATED
type DummygSupportedcsiRsDensity struct {
	Value utils.ENUMERATED
}

const (
	DummygSupportedcsiRsDensityEnumeratedNothing = iota
	DummygSupportedcsiRsDensityEnumeratedOne
	DummygSupportedcsiRsDensityEnumeratedThree
	DummygSupportedcsiRsDensityEnumeratedOneandthree
)
