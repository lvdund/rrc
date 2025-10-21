package ies

import "rrc/utils"

// NPDSCH-MultiTB-Config-NB-r16 ::= SEQUENCE
type NpdschMultitbConfigNbR16 struct {
	MultitbConfigR16   utils.ENUMERATED
	HarqAckbundlingR16 *utils.ENUMERATED
}
