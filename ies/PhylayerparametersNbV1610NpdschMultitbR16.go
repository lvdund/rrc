package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1610-npdsch-MultiTB-r16 ::= ENUMERATED
type PhylayerparametersNbV1610NpdschMultitbR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1610NpdschMultitbR16EnumeratedNothing = iota
	PhylayerparametersNbV1610NpdschMultitbR16EnumeratedSupported
)
