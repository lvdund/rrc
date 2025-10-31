package ies

import "rrc/utils"

// ServingCellConfig-enableTwoDefaultTCI-States-r16 ::= ENUMERATED
type ServingcellconfigEnabletwodefaulttciStatesR16 struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigEnabletwodefaulttciStatesR16EnumeratedNothing = iota
	ServingcellconfigEnabletwodefaulttciStatesR16EnumeratedEnabled
)
