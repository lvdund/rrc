package ies

import "rrc/utils"

// DeltaTxD-OffsetListPUCCH-r10 ::= SEQUENCE
// Extensible
type DeltatxdOffsetlistpucchR10 struct {
	DeltatxdOffsetpucchFormat1R10     utils.ENUMERATED
	DeltatxdOffsetpucchFormat1a1bR10  utils.ENUMERATED
	DeltatxdOffsetpucchFormat22a2bR10 utils.ENUMERATED
	DeltatxdOffsetpucchFormat3R10     utils.ENUMERATED
}
