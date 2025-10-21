package ies

import "rrc/utils"

// CE-PUSCH-MultiTB-Config-r16 ::= SEQUENCE
type CePuschMultitbConfigR16 struct {
	InterleavingR16 *utils.ENUMERATED
}
