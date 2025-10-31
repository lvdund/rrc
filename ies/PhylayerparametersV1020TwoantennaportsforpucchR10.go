package ies

import "rrc/utils"

// PhyLayerParameters-v1020-twoAntennaPortsForPUCCH-r10 ::= ENUMERATED
type PhylayerparametersV1020TwoantennaportsforpucchR10 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1020TwoantennaportsforpucchR10EnumeratedNothing = iota
	PhylayerparametersV1020TwoantennaportsforpucchR10EnumeratedSupported
)
