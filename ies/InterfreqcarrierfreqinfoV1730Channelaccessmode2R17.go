package ies

import "rrc/utils"

// InterFreqCarrierFreqInfo-v1730-channelAccessMode2-r17 ::= ENUMERATED
type InterfreqcarrierfreqinfoV1730Channelaccessmode2R17 struct {
	Value utils.ENUMERATED
}

const (
	InterfreqcarrierfreqinfoV1730Channelaccessmode2R17EnumeratedNothing = iota
	InterfreqcarrierfreqinfoV1730Channelaccessmode2R17EnumeratedEnabled
)
