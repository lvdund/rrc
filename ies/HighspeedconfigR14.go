package ies

import "rrc/utils"

// HighSpeedConfig-r14 ::= SEQUENCE
type HighspeedconfigR14 struct {
	HighspeedenhancedmeasflagR14         *utils.ENUMERATED
	HighspeedenhanceddemodulationflagR14 *utils.ENUMERATED
}
