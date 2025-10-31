package ies

import "rrc/utils"

// PhyLayerParameters-v1430-dmrs-LessUpPTS-r14 ::= ENUMERATED
type PhylayerparametersV1430DmrsLessupptsR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1430DmrsLessupptsR14EnumeratedNothing = iota
	PhylayerparametersV1430DmrsLessupptsR14EnumeratedSupported
)
