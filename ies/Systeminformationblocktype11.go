package ies

import "rrc/utils"

// SystemInformationBlockType11 ::= SEQUENCE
// Extensible
type Systeminformationblocktype11 struct {
	Messageidentifier           utils.BITSTRING
	Serialnumber                utils.BITSTRING
	Warningmessagesegmenttype   utils.ENUMERATED
	Warningmessagesegmentnumber utils.INTEGER
	Warningmessagesegment       utils.OCTETSTRING
	Datacodingscheme            *utils.OCTETSTRING
	Latenoncriticalextension    *utils.OCTETSTRING
}
