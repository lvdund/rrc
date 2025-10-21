package ies

import "rrc/utils"

// ULDedicatedMessageSegment-r16-IEs ::= SEQUENCE
type UldedicatedmessagesegmentR16Ies struct {
	SegmentnumberR16              utils.INTEGER
	RrcMessagesegmentcontainerR16 utils.OCTETSTRING
	RrcMessagesegmenttypeR16      utils.ENUMERATED
	Latenoncriticalextension      *utils.OCTETSTRING
	Noncriticalextension          *interface{}
}
