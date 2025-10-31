package ies

import "rrc/utils"

// CarrierInfoNR-r15-subcarrierSpacingSSB-r15 ::= ENUMERATED
type CarrierinfonrR15SubcarrierspacingssbR15 struct {
	Value utils.ENUMERATED
}

const (
	CarrierinfonrR15SubcarrierspacingssbR15EnumeratedNothing = iota
	CarrierinfonrR15SubcarrierspacingssbR15EnumeratedKhz15
	CarrierinfonrR15SubcarrierspacingssbR15EnumeratedKhz30
	CarrierinfonrR15SubcarrierspacingssbR15EnumeratedKhz120
	CarrierinfonrR15SubcarrierspacingssbR15EnumeratedKhz240
)
