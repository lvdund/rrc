package ies

import "rrc/utils"

// PhyLayerParameters-v1470-srs-UpPTS-6sym-r14 ::= ENUMERATED
type PhylayerparametersV1470SrsUppts6symR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1470SrsUppts6symR14EnumeratedNothing = iota
	PhylayerparametersV1470SrsUppts6symR14EnumeratedSupported
)
