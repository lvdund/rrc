package ies

import "rrc/utils"

// SystemInformationBlockType12-r9 ::= SEQUENCE
// Extensible
type Systeminformationblocktype12R9 struct {
	MessageidentifierR9           utils.BITSTRING `lb:16,ub:16`
	SerialnumberR9                utils.BITSTRING `lb:16,ub:16`
	WarningmessagesegmenttypeR9   Systeminformationblocktype12R9WarningmessagesegmenttypeR9
	WarningmessagesegmentnumberR9 utils.INTEGER `lb:0,ub:63`
	WarningmessagesegmentR9       utils.OCTETSTRING
	DatacodingschemeR9            *utils.OCTETSTRING `lb:1,ub:1`
	Latenoncriticalextension      *utils.OCTETSTRING
}
