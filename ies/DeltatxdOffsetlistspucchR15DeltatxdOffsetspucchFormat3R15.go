package ies

import "rrc/utils"

// DeltaTxD-OffsetListSPUCCH-r15-deltaTxD-OffsetSPUCCH-Format3-r15 ::= ENUMERATED
type DeltatxdOffsetlistspucchR15DeltatxdOffsetspucchFormat3R15 struct {
	Value utils.ENUMERATED
}

const (
	DeltatxdOffsetlistspucchR15DeltatxdOffsetspucchFormat3R15EnumeratedNothing = iota
	DeltatxdOffsetlistspucchR15DeltatxdOffsetspucchFormat3R15EnumeratedDb0
	DeltatxdOffsetlistspucchR15DeltatxdOffsetspucchFormat3R15EnumeratedDb_2
)
