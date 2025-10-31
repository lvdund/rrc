package ies

import "rrc/utils"

// SRS-CapabilityPerBandPair-v14b0-srs-FlexibleTiming-r14 ::= ENUMERATED
type SrsCapabilityperbandpairV14b0SrsFlexibletimingR14 struct {
	Value utils.ENUMERATED
}

const (
	SrsCapabilityperbandpairV14b0SrsFlexibletimingR14EnumeratedNothing = iota
	SrsCapabilityperbandpairV14b0SrsFlexibletimingR14EnumeratedSupported
)
