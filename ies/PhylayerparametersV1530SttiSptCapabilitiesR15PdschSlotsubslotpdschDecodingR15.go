package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-pdsch-SlotSubslotPDSCH-Decoding-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15PdschSlotsubslotpdschDecodingR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15PdschSlotsubslotpdschDecodingR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15PdschSlotsubslotpdschDecodingR15EnumeratedSupported
)
