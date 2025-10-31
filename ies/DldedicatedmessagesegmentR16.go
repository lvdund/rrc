package ies

import "rrc/utils"

// DLDedicatedMessageSegment-r16-IEs ::= SEQUENCE
type DldedicatedmessagesegmentR16 struct {
	SegmentnumberR16              utils.INTEGER `lb:0,ub:4`
	RrcMessagesegmentcontainerR16 utils.OCTETSTRING
	RrcMessagesegmenttypeR16      DldedicatedmessagesegmentR16IesRrcMessagesegmenttypeR16
	Latenoncriticalextension      *utils.OCTETSTRING
	Noncriticalextension          *DldedicatedmessagesegmentR16IesNoncriticalextension
}
