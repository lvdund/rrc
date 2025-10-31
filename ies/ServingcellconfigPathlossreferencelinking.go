package ies

import "rrc/utils"

// ServingCellConfig-pathlossReferenceLinking ::= ENUMERATED
type ServingcellconfigPathlossreferencelinking struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigPathlossreferencelinkingEnumeratedNothing = iota
	ServingcellconfigPathlossreferencelinkingEnumeratedSpcell
	ServingcellconfigPathlossreferencelinkingEnumeratedScell
)
