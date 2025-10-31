package ies

import "rrc/utils"

// PhyLayerParameters-v9d0-tm5-FDD-r9 ::= ENUMERATED
type PhylayerparametersV9d0Tm5FddR9 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV9d0Tm5FddR9EnumeratedNothing = iota
	PhylayerparametersV9d0Tm5FddR9EnumeratedSupported
)
