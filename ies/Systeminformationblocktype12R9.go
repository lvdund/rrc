package ies

import "rrc/utils"

// SystemInformationBlockType12-r9 ::= SEQUENCE
// Extensible
type Systeminformationblocktype12R9 struct {
	MessageidentifierR9           utils.BITSTRING
	SerialnumberR9                utils.BITSTRING
	WarningmessagesegmenttypeR9   utils.ENUMERATED
	WarningmessagesegmentnumberR9 utils.INTEGER
	WarningmessagesegmentR9       utils.OCTETSTRING
	DatacodingschemeR9            *utils.OCTETSTRING
	Latenoncriticalextension      *utils.OCTETSTRING
}
