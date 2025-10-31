package ies

import "rrc/utils"

// PhyLayerParameters-v1130-ul-CoMP-r11 ::= ENUMERATED
type PhylayerparametersV1130UlCompR11 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1130UlCompR11EnumeratedNothing = iota
	PhylayerparametersV1130UlCompR11EnumeratedSupported
)
