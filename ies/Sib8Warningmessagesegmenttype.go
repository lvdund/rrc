package ies

import "rrc/utils"

// SIB8-warningMessageSegmentType ::= ENUMERATED
type Sib8Warningmessagesegmenttype struct {
	Value utils.ENUMERATED
}

const (
	Sib8WarningmessagesegmenttypeEnumeratedNothing = iota
	Sib8WarningmessagesegmenttypeEnumeratedNotlastsegment
	Sib8WarningmessagesegmenttypeEnumeratedLastsegment
)
