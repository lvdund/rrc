package ies

import "rrc/utils"

// PhyLayerParameters-v1540-crs-IM-TM1-toTM9-OneRX-Port-v1540 ::= ENUMERATED
type PhylayerparametersV1540CrsImTm1Totm9OnerxPortV1540 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1540CrsImTm1Totm9OnerxPortV1540EnumeratedNothing = iota
	PhylayerparametersV1540CrsImTm1Totm9OnerxPortV1540EnumeratedSupported
)
