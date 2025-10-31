package ies

import "rrc/utils"

// PhyLayerParameters-v1020-crossCarrierScheduling-r10 ::= ENUMERATED
type PhylayerparametersV1020CrosscarrierschedulingR10 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1020CrosscarrierschedulingR10EnumeratedNothing = iota
	PhylayerparametersV1020CrosscarrierschedulingR10EnumeratedSupported
)
