package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1430-multiCarrier-NPRACH-r14 ::= ENUMERATED
type PhylayerparametersNbV1430MulticarrierNprachR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1430MulticarrierNprachR14EnumeratedNothing = iota
	PhylayerparametersNbV1430MulticarrierNprachR14EnumeratedSupported
)
