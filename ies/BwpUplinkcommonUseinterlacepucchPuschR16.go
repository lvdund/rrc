package ies

import "rrc/utils"

// BWP-UplinkCommon-useInterlacePUCCH-PUSCH-r16 ::= ENUMERATED
type BwpUplinkcommonUseinterlacepucchPuschR16 struct {
	Value utils.ENUMERATED
}

const (
	BwpUplinkcommonUseinterlacepucchPuschR16EnumeratedNothing = iota
	BwpUplinkcommonUseinterlacepucchPuschR16EnumeratedEnabled
)
