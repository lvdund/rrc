package ies

import "rrc/utils"

// CE-PDSCH-MultiTB-Config-r16 ::= SEQUENCE
type CePdschMultitbConfigR16 struct {
	InterleavingR16    *utils.ENUMERATED
	HarqAckbundlingR16 *utils.ENUMERATED
}
