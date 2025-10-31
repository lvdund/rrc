package ies

import "rrc/utils"

// PhyLayerParameters-v1020-pmi-Disabling-r10 ::= ENUMERATED
type PhylayerparametersV1020PmiDisablingR10 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1020PmiDisablingR10EnumeratedNothing = iota
	PhylayerparametersV1020PmiDisablingR10EnumeratedSupported
)
