package ies

import "rrc/utils"

// PhyLayerParameters-v1530-crs-IntfMitig-r15 ::= ENUMERATED
type PhylayerparametersV1530CrsIntfmitigR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530CrsIntfmitigR15EnumeratedNothing = iota
	PhylayerparametersV1530CrsIntfmitigR15EnumeratedSupported
)
