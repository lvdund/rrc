package ies

import "rrc/utils"

// ServingCellConfig-dummy1 ::= ENUMERATED
type ServingcellconfigDummy1 struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigDummy1EnumeratedNothing = iota
	ServingcellconfigDummy1EnumeratedEnabled
)
