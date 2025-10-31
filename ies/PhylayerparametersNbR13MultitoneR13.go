package ies

import "rrc/utils"

// PhyLayerParameters-NB-r13-multiTone-r13 ::= ENUMERATED
type PhylayerparametersNbR13MultitoneR13 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbR13MultitoneR13EnumeratedNothing = iota
	PhylayerparametersNbR13MultitoneR13EnumeratedSupported
)
