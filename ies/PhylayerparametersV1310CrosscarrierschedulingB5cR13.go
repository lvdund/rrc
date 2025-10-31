package ies

import "rrc/utils"

// PhyLayerParameters-v1310-crossCarrierScheduling-B5C-r13 ::= ENUMERATED
type PhylayerparametersV1310CrosscarrierschedulingB5cR13 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1310CrosscarrierschedulingB5cR13EnumeratedNothing = iota
	PhylayerparametersV1310CrosscarrierschedulingB5cR13EnumeratedSupported
)
