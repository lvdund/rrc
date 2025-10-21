package ies

import "rrc/utils"

// DeltaTxD-OffsetListSPUCCH-r15 ::= SEQUENCE
// Extensible
type DeltatxdOffsetlistspucchR15 struct {
	DeltatxdOffsetspucchFormat1R15  utils.ENUMERATED
	DeltatxdOffsetspucchFormat1aR15 utils.ENUMERATED
	DeltatxdOffsetspucchFormat1bR15 utils.ENUMERATED
	DeltatxdOffsetspucchFormat3R15  utils.ENUMERATED
}
