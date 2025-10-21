package ies

import "rrc/utils"

// SystemInformationBlockType10 ::= SEQUENCE
// Extensible
type Systeminformationblocktype10 struct {
	Messageidentifier        utils.BITSTRING
	Serialnumber             utils.BITSTRING
	Warningtype              utils.OCTETSTRING
	Dummy                    *utils.OCTETSTRING
	Latenoncriticalextension *utils.OCTETSTRING
}
