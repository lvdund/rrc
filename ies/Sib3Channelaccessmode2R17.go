package ies

import "rrc/utils"

// SIB3-channelAccessMode2-r17 ::= ENUMERATED
type Sib3Channelaccessmode2R17 struct {
	Value utils.ENUMERATED
}

const (
	Sib3Channelaccessmode2R17EnumeratedNothing = iota
	Sib3Channelaccessmode2R17EnumeratedEnabled
)
