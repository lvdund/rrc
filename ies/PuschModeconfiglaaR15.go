package ies

import "rrc/utils"

// PUSCH-ModeConfigLAA-r15 ::= SEQUENCE
type PuschModeconfiglaaR15 struct {
	LaaPuschMode1 utils.BOOLEAN
	LaaPuschMode2 utils.BOOLEAN
	LaaPuschMode3 utils.BOOLEAN
}
