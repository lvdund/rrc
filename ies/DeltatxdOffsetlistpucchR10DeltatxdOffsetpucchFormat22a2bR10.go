package ies

import "rrc/utils"

// DeltaTxD-OffsetListPUCCH-r10-deltaTxD-OffsetPUCCH-Format22a2b-r10 ::= ENUMERATED
type DeltatxdOffsetlistpucchR10DeltatxdOffsetpucchFormat22a2bR10 struct {
	Value utils.ENUMERATED
}

const (
	DeltatxdOffsetlistpucchR10DeltatxdOffsetpucchFormat22a2bR10EnumeratedNothing = iota
	DeltatxdOffsetlistpucchR10DeltatxdOffsetpucchFormat22a2bR10EnumeratedDb0
	DeltatxdOffsetlistpucchR10DeltatxdOffsetpucchFormat22a2bR10EnumeratedDb_2
)
