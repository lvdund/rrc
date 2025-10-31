package ies

import "rrc/utils"

// SIB6 ::= SEQUENCE
// Extensible
type Sib6 struct {
	Messageidentifier        utils.BITSTRING   `lb:16,ub:16`
	Serialnumber             utils.BITSTRING   `lb:16,ub:16`
	Warningtype              utils.OCTETSTRING `lb:2,ub:2`
	Latenoncriticalextension *utils.OCTETSTRING
}
