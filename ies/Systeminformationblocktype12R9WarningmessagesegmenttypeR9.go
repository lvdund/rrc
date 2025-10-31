package ies

import "rrc/utils"

// SystemInformationBlockType12-r9-warningMessageSegmentType-r9 ::= ENUMERATED
type Systeminformationblocktype12R9WarningmessagesegmenttypeR9 struct {
	Value utils.ENUMERATED
}

const (
	Systeminformationblocktype12R9WarningmessagesegmenttypeR9EnumeratedNothing = iota
	Systeminformationblocktype12R9WarningmessagesegmenttypeR9EnumeratedNotlastsegment
	Systeminformationblocktype12R9WarningmessagesegmenttypeR9EnumeratedLastsegment
)
