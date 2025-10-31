package ies

import "rrc/utils"

// CarrierFreqNR-r15-subcarrierSpacingSSB-r15 ::= ENUMERATED
type CarrierfreqnrR15SubcarrierspacingssbR15 struct {
	Value utils.ENUMERATED
}

const (
	CarrierfreqnrR15SubcarrierspacingssbR15EnumeratedNothing = iota
	CarrierfreqnrR15SubcarrierspacingssbR15EnumeratedKhz15
	CarrierfreqnrR15SubcarrierspacingssbR15EnumeratedKhz30
	CarrierfreqnrR15SubcarrierspacingssbR15EnumeratedKhz120
	CarrierfreqnrR15SubcarrierspacingssbR15EnumeratedKhz240
)
