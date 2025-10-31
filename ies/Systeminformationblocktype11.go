package ies

import "rrc/utils"

// SystemInformationBlockType11 ::= SEQUENCE
// Extensible
type Systeminformationblocktype11 struct {
	Messageidentifier           utils.BITSTRING `lb:16,ub:16`
	Serialnumber                utils.BITSTRING `lb:16,ub:16`
	Warningmessagesegmenttype   Systeminformationblocktype11Warningmessagesegmenttype
	Warningmessagesegmentnumber utils.INTEGER `lb:0,ub:63`
	Warningmessagesegment       utils.OCTETSTRING
	Datacodingscheme            *utils.OCTETSTRING `lb:1,ub:1`
	Latenoncriticalextension    *utils.OCTETSTRING
}
