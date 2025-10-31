package ies

import "rrc/utils"

// MIMO-LayersUL ::= ENUMERATED
type MimoLayersul struct {
	Value utils.ENUMERATED
}

const (
	MimoLayersulEnumeratedNothing = iota
	MimoLayersulEnumeratedOnelayer
	MimoLayersulEnumeratedTwolayers
	MimoLayersulEnumeratedFourlayers
)
