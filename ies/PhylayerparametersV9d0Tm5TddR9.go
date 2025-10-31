package ies

import "rrc/utils"

// PhyLayerParameters-v9d0-tm5-TDD-r9 ::= ENUMERATED
type PhylayerparametersV9d0Tm5TddR9 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV9d0Tm5TddR9EnumeratedNothing = iota
	PhylayerparametersV9d0Tm5TddR9EnumeratedSupported
)
