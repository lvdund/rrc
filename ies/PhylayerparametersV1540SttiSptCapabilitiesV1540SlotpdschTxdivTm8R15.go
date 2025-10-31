package ies

import "rrc/utils"

// PhyLayerParameters-v1540-stti-SPT-Capabilities-v1540-slotPDSCH-TxDiv-TM8-r15 ::= ENUMERATED
type PhylayerparametersV1540SttiSptCapabilitiesV1540SlotpdschTxdivTm8R15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1540SttiSptCapabilitiesV1540SlotpdschTxdivTm8R15EnumeratedNothing = iota
	PhylayerparametersV1540SttiSptCapabilitiesV1540SlotpdschTxdivTm8R15EnumeratedSupported
)
