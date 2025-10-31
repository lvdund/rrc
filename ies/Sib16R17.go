package ies

import "rrc/utils"

// SIB16-r17 ::= SEQUENCE
// Extensible
type Sib16R17 struct {
	FreqprioritylistslicingR17 *FreqprioritylistslicingR17
	Latenoncriticalextension   *utils.OCTETSTRING
}
