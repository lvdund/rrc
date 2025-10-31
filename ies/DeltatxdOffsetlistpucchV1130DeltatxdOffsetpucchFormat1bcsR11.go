package ies

import "rrc/utils"

// DeltaTxD-OffsetListPUCCH-v1130-deltaTxD-OffsetPUCCH-Format1bCS-r11 ::= ENUMERATED
type DeltatxdOffsetlistpucchV1130DeltatxdOffsetpucchFormat1bcsR11 struct {
	Value utils.ENUMERATED
}

const (
	DeltatxdOffsetlistpucchV1130DeltatxdOffsetpucchFormat1bcsR11EnumeratedNothing = iota
	DeltatxdOffsetlistpucchV1130DeltatxdOffsetpucchFormat1bcsR11EnumeratedDb0
	DeltatxdOffsetlistpucchV1130DeltatxdOffsetpucchFormat1bcsR11EnumeratedDb_1
)
