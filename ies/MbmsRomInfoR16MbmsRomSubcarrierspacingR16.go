package ies

import "rrc/utils"

// MBMS-ROM-Info-r16-mbms-ROM-SubcarrierSpacing-r16 ::= ENUMERATED
type MbmsRomInfoR16MbmsRomSubcarrierspacingR16 struct {
	Value utils.ENUMERATED
}

const (
	MbmsRomInfoR16MbmsRomSubcarrierspacingR16EnumeratedNothing = iota
	MbmsRomInfoR16MbmsRomSubcarrierspacingR16EnumeratedKhz2dot5
	MbmsRomInfoR16MbmsRomSubcarrierspacingR16EnumeratedKhz0dot37
)
