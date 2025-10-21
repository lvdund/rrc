package ies

import "rrc/utils"

// SRS-CapabilityPerBandPair-v14b0 ::= SEQUENCE
type SrsCapabilityperbandpairV14b0 struct {
	SrsFlexibletimingR14      *utils.ENUMERATED
	SrsHarqReferenceconfigR14 *utils.ENUMERATED
}
