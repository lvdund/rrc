package ies

import "rrc/utils"

// PhyLayerParameters-v1020-multiClusterPUSCH-WithinCC-r10 ::= ENUMERATED
type PhylayerparametersV1020MulticlusterpuschWithinccR10 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1020MulticlusterpuschWithinccR10EnumeratedNothing = iota
	PhylayerparametersV1020MulticlusterpuschWithinccR10EnumeratedSupported
)
