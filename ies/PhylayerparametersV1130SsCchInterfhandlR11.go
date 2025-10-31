package ies

import "rrc/utils"

// PhyLayerParameters-v1130-ss-CCH-InterfHandl-r11 ::= ENUMERATED
type PhylayerparametersV1130SsCchInterfhandlR11 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1130SsCchInterfhandlR11EnumeratedNothing = iota
	PhylayerparametersV1130SsCchInterfhandlR11EnumeratedSupported
)
