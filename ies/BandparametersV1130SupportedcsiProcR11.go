package ies

import "rrc/utils"

// BandParameters-v1130-supportedCSI-Proc-r11 ::= ENUMERATED
type BandparametersV1130SupportedcsiProcR11 struct {
	Value utils.ENUMERATED
}

const (
	BandparametersV1130SupportedcsiProcR11EnumeratedNothing = iota
	BandparametersV1130SupportedcsiProcR11EnumeratedN1
	BandparametersV1130SupportedcsiProcR11EnumeratedN3
	BandparametersV1130SupportedcsiProcR11EnumeratedN4
)
