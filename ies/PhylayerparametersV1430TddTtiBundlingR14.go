package ies

import "rrc/utils"

// PhyLayerParameters-v1430-tdd-TTI-Bundling-r14 ::= ENUMERATED
type PhylayerparametersV1430TddTtiBundlingR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1430TddTtiBundlingR14EnumeratedNothing = iota
	PhylayerparametersV1430TddTtiBundlingR14EnumeratedSupported
)
