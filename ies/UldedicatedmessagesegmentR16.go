package ies

import "rrc/utils"

// ULDedicatedMessageSegment-r16-IEs ::= SEQUENCE
type UldedicatedmessagesegmentR16 struct {
	SegmentnumberR16              utils.INTEGER `lb:0,ub:15`
	RrcMessagesegmentcontainerR16 utils.OCTETSTRING
	RrcMessagesegmenttypeR16      UldedicatedmessagesegmentR16IesRrcMessagesegmenttypeR16
	Latenoncriticalextension      *utils.OCTETSTRING
	Noncriticalextension          *UldedicatedmessagesegmentR16IesNoncriticalextension
}
