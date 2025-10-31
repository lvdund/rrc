package ies

import "rrc/utils"

// SRS-CapabilityPerBandPair-v1610-addSRS-CarrierSwitching-r16 ::= ENUMERATED
type SrsCapabilityperbandpairV1610AddsrsCarrierswitchingR16 struct {
	Value utils.ENUMERATED
}

const (
	SrsCapabilityperbandpairV1610AddsrsCarrierswitchingR16EnumeratedNothing = iota
	SrsCapabilityperbandpairV1610AddsrsCarrierswitchingR16EnumeratedSupported
)
