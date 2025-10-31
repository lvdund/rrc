package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1610-npusch-MultiTB-r16 ::= ENUMERATED
type PhylayerparametersNbV1610NpuschMultitbR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1610NpuschMultitbR16EnumeratedNothing = iota
	PhylayerparametersNbV1610NpuschMultitbR16EnumeratedSupported
)
