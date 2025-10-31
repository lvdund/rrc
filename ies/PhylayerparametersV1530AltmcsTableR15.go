package ies

import "rrc/utils"

// PhyLayerParameters-v1530-altMCS-Table-r15 ::= ENUMERATED
type PhylayerparametersV1530AltmcsTableR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530AltmcsTableR15EnumeratedNothing = iota
	PhylayerparametersV1530AltmcsTableR15EnumeratedSupported
)
