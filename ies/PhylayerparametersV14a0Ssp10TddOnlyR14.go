package ies

import "rrc/utils"

// PhyLayerParameters-v14a0-ssp10-TDD-Only-r14 ::= ENUMERATED
type PhylayerparametersV14a0Ssp10TddOnlyR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV14a0Ssp10TddOnlyR14EnumeratedNothing = iota
	PhylayerparametersV14a0Ssp10TddOnlyR14EnumeratedSupported
)
