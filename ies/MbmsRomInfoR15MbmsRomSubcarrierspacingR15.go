package ies

import "rrc/utils"

// MBMS-ROM-Info-r15-mbms-ROM-SubcarrierSpacing-r15 ::= ENUMERATED
type MbmsRomInfoR15MbmsRomSubcarrierspacingR15 struct {
	Value utils.ENUMERATED
}

const (
	MbmsRomInfoR15MbmsRomSubcarrierspacingR15EnumeratedNothing = iota
	MbmsRomInfoR15MbmsRomSubcarrierspacingR15EnumeratedKhz15
	MbmsRomInfoR15MbmsRomSubcarrierspacingR15EnumeratedKhz7dot5
	MbmsRomInfoR15MbmsRomSubcarrierspacingR15EnumeratedKhz1dot25
)
