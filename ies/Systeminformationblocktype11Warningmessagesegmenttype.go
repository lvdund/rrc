package ies

import "rrc/utils"

// SystemInformationBlockType11-warningMessageSegmentType ::= ENUMERATED
type Systeminformationblocktype11Warningmessagesegmenttype struct {
	Value utils.ENUMERATED
}

const (
	Systeminformationblocktype11WarningmessagesegmenttypeEnumeratedNothing = iota
	Systeminformationblocktype11WarningmessagesegmenttypeEnumeratedNotlastsegment
	Systeminformationblocktype11WarningmessagesegmenttypeEnumeratedLastsegment
)
