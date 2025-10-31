package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-slotPDSCH-TxDiv-TM9and10 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15SlotpdschTxdivTm9and10 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15SlotpdschTxdivTm9and10EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15SlotpdschTxdivTm9and10EnumeratedSupported
)
