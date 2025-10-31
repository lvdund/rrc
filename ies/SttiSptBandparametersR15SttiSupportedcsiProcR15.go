package ies

import "rrc/utils"

// STTI-SPT-BandParameters-r15-sTTI-SupportedCSI-Proc-r15 ::= ENUMERATED
type SttiSptBandparametersR15SttiSupportedcsiProcR15 struct {
	Value utils.ENUMERATED
}

const (
	SttiSptBandparametersR15SttiSupportedcsiProcR15EnumeratedNothing = iota
	SttiSptBandparametersR15SttiSupportedcsiProcR15EnumeratedN1
	SttiSptBandparametersR15SttiSupportedcsiProcR15EnumeratedN3
	SttiSptBandparametersR15SttiSupportedcsiProcR15EnumeratedN4
)
