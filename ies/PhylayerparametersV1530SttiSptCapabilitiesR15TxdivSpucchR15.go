package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-txDiv-SPUCCH-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15TxdivSpucchR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15TxdivSpucchR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15TxdivSpucchR15EnumeratedSupported
)
