package ies

import "rrc/utils"

// SIB7-warningMessageSegmentType ::= ENUMERATED
type Sib7Warningmessagesegmenttype struct {
	Value utils.ENUMERATED
}

const (
	Sib7WarningmessagesegmenttypeEnumeratedNothing = iota
	Sib7WarningmessagesegmenttypeEnumeratedNotlastsegment
	Sib7WarningmessagesegmenttypeEnumeratedLastsegment
)
