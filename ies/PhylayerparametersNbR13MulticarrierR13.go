package ies

import "rrc/utils"

// PhyLayerParameters-NB-r13-multiCarrier-r13 ::= ENUMERATED
type PhylayerparametersNbR13MulticarrierR13 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbR13MulticarrierR13EnumeratedNothing = iota
	PhylayerparametersNbR13MulticarrierR13EnumeratedSupported
)
