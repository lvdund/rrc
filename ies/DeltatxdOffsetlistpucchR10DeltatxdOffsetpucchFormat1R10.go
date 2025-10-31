package ies

import "rrc/utils"

// DeltaTxD-OffsetListPUCCH-r10-deltaTxD-OffsetPUCCH-Format1-r10 ::= ENUMERATED
type DeltatxdOffsetlistpucchR10DeltatxdOffsetpucchFormat1R10 struct {
	Value utils.ENUMERATED
}

const (
	DeltatxdOffsetlistpucchR10DeltatxdOffsetpucchFormat1R10EnumeratedNothing = iota
	DeltatxdOffsetlistpucchR10DeltatxdOffsetpucchFormat1R10EnumeratedDb0
	DeltatxdOffsetlistpucchR10DeltatxdOffsetpucchFormat1R10EnumeratedDb_2
)
