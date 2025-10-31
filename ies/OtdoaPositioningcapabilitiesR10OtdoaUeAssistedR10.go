package ies

import "rrc/utils"

// OTDOA-PositioningCapabilities-r10-otdoa-UE-Assisted-r10 ::= ENUMERATED
type OtdoaPositioningcapabilitiesR10OtdoaUeAssistedR10 struct {
	Value utils.ENUMERATED
}

const (
	OtdoaPositioningcapabilitiesR10OtdoaUeAssistedR10EnumeratedNothing = iota
	OtdoaPositioningcapabilitiesR10OtdoaUeAssistedR10EnumeratedSupported
)
