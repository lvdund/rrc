package ies

import "rrc/utils"

// DLDedicatedMessageSegment-r16-IEs-rrc-MessageSegmentType-r16 ::= ENUMERATED
type DldedicatedmessagesegmentR16IesRrcMessagesegmenttypeR16 struct {
	Value utils.ENUMERATED
}

const (
	DldedicatedmessagesegmentR16IesRrcMessagesegmenttypeR16EnumeratedNothing = iota
	DldedicatedmessagesegmentR16IesRrcMessagesegmenttypeR16EnumeratedNotlastsegment
	DldedicatedmessagesegmentR16IesRrcMessagesegmenttypeR16EnumeratedLastsegment
)
