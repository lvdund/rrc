package ies

import "rrc/utils"

// SystemInformationBlockType10 ::= SEQUENCE
// Extensible
type Systeminformationblocktype10 struct {
	Messageidentifier        utils.BITSTRING    `lb:16,ub:16`
	Serialnumber             utils.BITSTRING    `lb:16,ub:16`
	Warningtype              utils.OCTETSTRING  `lb:2,ub:2`
	Dummy                    *utils.OCTETSTRING `lb:50,ub:50`
	Latenoncriticalextension *utils.OCTETSTRING
}
