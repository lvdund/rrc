package ies

import "rrc/utils"

// BandParameters-r11-supportedCSI-Proc-r11 ::= ENUMERATED
type BandparametersR11SupportedcsiProcR11 struct {
	Value utils.ENUMERATED
}

const (
	BandparametersR11SupportedcsiProcR11EnumeratedNothing = iota
	BandparametersR11SupportedcsiProcR11EnumeratedN1
	BandparametersR11SupportedcsiProcR11EnumeratedN3
	BandparametersR11SupportedcsiProcR11EnumeratedN4
)
