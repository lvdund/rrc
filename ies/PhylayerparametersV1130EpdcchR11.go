package ies

import "rrc/utils"

// PhyLayerParameters-v1130-ePDCCH-r11 ::= ENUMERATED
type PhylayerparametersV1130EpdcchR11 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1130EpdcchR11EnumeratedNothing = iota
	PhylayerparametersV1130EpdcchR11EnumeratedSupported
)
