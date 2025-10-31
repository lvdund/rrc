package ies

import "rrc/utils"

// BWP-UplinkDedicated-useInterlacePUCCH-PUSCH-r16 ::= ENUMERATED
type BwpUplinkdedicatedUseinterlacepucchPuschR16 struct {
	Value utils.ENUMERATED
}

const (
	BwpUplinkdedicatedUseinterlacepucchPuschR16EnumeratedNothing = iota
	BwpUplinkdedicatedUseinterlacepucchPuschR16EnumeratedEnabled
)
