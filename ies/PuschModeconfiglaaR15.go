package ies

import "rrc/utils"

// PUSCH-ModeConfigLAA-r15 ::= SEQUENCE
type PuschModeconfiglaaR15 struct {
	LaaPuschMode1 bool
	LaaPuschMode2 bool
	LaaPuschMode3 bool
}
