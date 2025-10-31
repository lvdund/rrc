package ies

import "rrc/utils"

// SRS-CapabilityPerBandPair-v14b0-srs-HARQ-ReferenceConfig-r14 ::= ENUMERATED
type SrsCapabilityperbandpairV14b0SrsHarqReferenceconfigR14 struct {
	Value utils.ENUMERATED
}

const (
	SrsCapabilityperbandpairV14b0SrsHarqReferenceconfigR14EnumeratedNothing = iota
	SrsCapabilityperbandpairV14b0SrsHarqReferenceconfigR14EnumeratedSupported
)
