package ies

import "rrc/utils"

// RS-ConfigSSB-NR-r15-subcarrierSpacingSSB-r15 ::= ENUMERATED
type RsConfigssbNrR15SubcarrierspacingssbR15 struct {
	Value utils.ENUMERATED
}

const (
	RsConfigssbNrR15SubcarrierspacingssbR15EnumeratedNothing = iota
	RsConfigssbNrR15SubcarrierspacingssbR15EnumeratedKhz15
	RsConfigssbNrR15SubcarrierspacingssbR15EnumeratedKhz30
	RsConfigssbNrR15SubcarrierspacingssbR15EnumeratedKhz120
	RsConfigssbNrR15SubcarrierspacingssbR15EnumeratedKhz240
)
