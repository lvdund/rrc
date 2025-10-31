package ies

import "rrc/utils"

// DeltaTxD-OffsetListSPUCCH-r15-deltaTxD-OffsetSPUCCH-Format1b-r15 ::= ENUMERATED
type DeltatxdOffsetlistspucchR15DeltatxdOffsetspucchFormat1bR15 struct {
	Value utils.ENUMERATED
}

const (
	DeltatxdOffsetlistspucchR15DeltatxdOffsetspucchFormat1bR15EnumeratedNothing = iota
	DeltatxdOffsetlistspucchR15DeltatxdOffsetspucchFormat1bR15EnumeratedDb0
	DeltatxdOffsetlistspucchR15DeltatxdOffsetspucchFormat1bR15EnumeratedDb_2
)
