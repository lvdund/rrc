package ies

import "rrc/utils"

// DeltaTxD-OffsetListPUCCH-r10-deltaTxD-OffsetPUCCH-Format3-r10 ::= ENUMERATED
type DeltatxdOffsetlistpucchR10DeltatxdOffsetpucchFormat3R10 struct {
	Value utils.ENUMERATED
}

const (
	DeltatxdOffsetlistpucchR10DeltatxdOffsetpucchFormat3R10EnumeratedNothing = iota
	DeltatxdOffsetlistpucchR10DeltatxdOffsetpucchFormat3R10EnumeratedDb0
	DeltatxdOffsetlistpucchR10DeltatxdOffsetpucchFormat3R10EnumeratedDb_2
)
