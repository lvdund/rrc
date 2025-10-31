package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1610-npusch-MultiTB-Interleaving-r16 ::= ENUMERATED
type PhylayerparametersNbV1610NpuschMultitbInterleavingR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1610NpuschMultitbInterleavingR16EnumeratedNothing = iota
	PhylayerparametersNbV1610NpuschMultitbInterleavingR16EnumeratedSupported
)
