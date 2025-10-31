package ies

import "rrc/utils"

// PhyLayerParameters-v1130-crs-InterfHandl-r11 ::= ENUMERATED
type PhylayerparametersV1130CrsInterfhandlR11 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1130CrsInterfhandlR11EnumeratedNothing = iota
	PhylayerparametersV1130CrsInterfhandlR11EnumeratedSupported
)
