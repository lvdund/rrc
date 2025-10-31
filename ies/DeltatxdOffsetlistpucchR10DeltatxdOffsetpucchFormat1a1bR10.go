package ies

import "rrc/utils"

// DeltaTxD-OffsetListPUCCH-r10-deltaTxD-OffsetPUCCH-Format1a1b-r10 ::= ENUMERATED
type DeltatxdOffsetlistpucchR10DeltatxdOffsetpucchFormat1a1bR10 struct {
	Value utils.ENUMERATED
}

const (
	DeltatxdOffsetlistpucchR10DeltatxdOffsetpucchFormat1a1bR10EnumeratedNothing = iota
	DeltatxdOffsetlistpucchR10DeltatxdOffsetpucchFormat1a1bR10EnumeratedDb0
	DeltatxdOffsetlistpucchR10DeltatxdOffsetpucchFormat1a1bR10EnumeratedDb_2
)
