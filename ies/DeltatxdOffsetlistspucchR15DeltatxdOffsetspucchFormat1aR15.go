package ies

import "rrc/utils"

// DeltaTxD-OffsetListSPUCCH-r15-deltaTxD-OffsetSPUCCH-Format1a-r15 ::= ENUMERATED
type DeltatxdOffsetlistspucchR15DeltatxdOffsetspucchFormat1aR15 struct {
	Value utils.ENUMERATED
}

const (
	DeltatxdOffsetlistspucchR15DeltatxdOffsetspucchFormat1aR15EnumeratedNothing = iota
	DeltatxdOffsetlistspucchR15DeltatxdOffsetspucchFormat1aR15EnumeratedDb0
	DeltatxdOffsetlistspucchR15DeltatxdOffsetspucchFormat1aR15EnumeratedDb_2
)
