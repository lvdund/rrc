package ies

import "rrc/utils"

// DeltaTxD-OffsetListSPUCCH-r15-deltaTxD-OffsetSPUCCH-Format1-r15 ::= ENUMERATED
type DeltatxdOffsetlistspucchR15DeltatxdOffsetspucchFormat1R15 struct {
	Value utils.ENUMERATED
}

const (
	DeltatxdOffsetlistspucchR15DeltatxdOffsetspucchFormat1R15EnumeratedNothing = iota
	DeltatxdOffsetlistspucchR15DeltatxdOffsetspucchFormat1R15EnumeratedDb0
	DeltatxdOffsetlistspucchR15DeltatxdOffsetspucchFormat1R15EnumeratedDb_2
)
