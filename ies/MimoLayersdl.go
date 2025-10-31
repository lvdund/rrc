package ies

import "rrc/utils"

// MIMO-LayersDL ::= ENUMERATED
type MimoLayersdl struct {
	Value utils.ENUMERATED
}

const (
	MimoLayersdlEnumeratedNothing = iota
	MimoLayersdlEnumeratedTwolayers
	MimoLayersdlEnumeratedFourlayers
	MimoLayersdlEnumeratedEightlayers
)
