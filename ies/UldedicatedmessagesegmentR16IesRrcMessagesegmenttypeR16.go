package ies

import "rrc/utils"

// ULDedicatedMessageSegment-r16-IEs-rrc-MessageSegmentType-r16 ::= ENUMERATED
type UldedicatedmessagesegmentR16IesRrcMessagesegmenttypeR16 struct {
	Value utils.ENUMERATED
}

const (
	UldedicatedmessagesegmentR16IesRrcMessagesegmenttypeR16EnumeratedNothing = iota
	UldedicatedmessagesegmentR16IesRrcMessagesegmenttypeR16EnumeratedNotlastsegment
	UldedicatedmessagesegmentR16IesRrcMessagesegmenttypeR16EnumeratedLastsegment
)
