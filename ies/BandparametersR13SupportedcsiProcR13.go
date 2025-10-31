package ies

import "rrc/utils"

// BandParameters-r13-supportedCSI-Proc-r13 ::= ENUMERATED
type BandparametersR13SupportedcsiProcR13 struct {
	Value utils.ENUMERATED
}

const (
	BandparametersR13SupportedcsiProcR13EnumeratedNothing = iota
	BandparametersR13SupportedcsiProcR13EnumeratedN1
	BandparametersR13SupportedcsiProcR13EnumeratedN3
	BandparametersR13SupportedcsiProcR13EnumeratedN4
)
