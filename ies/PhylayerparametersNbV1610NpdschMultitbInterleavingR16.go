package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1610-npdsch-MultiTB-Interleaving-r16 ::= ENUMERATED
type PhylayerparametersNbV1610NpdschMultitbInterleavingR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1610NpdschMultitbInterleavingR16EnumeratedNothing = iota
	PhylayerparametersNbV1610NpdschMultitbInterleavingR16EnumeratedSupported
)
