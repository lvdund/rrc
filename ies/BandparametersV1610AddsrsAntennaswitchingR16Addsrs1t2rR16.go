package ies

import "rrc/utils"

// BandParameters-v1610-addSRS-AntennaSwitching-r16-addSRS-1T2R-r16 ::= ENUMERATED
type BandparametersV1610AddsrsAntennaswitchingR16Addsrs1t2rR16 struct {
	Value utils.ENUMERATED
}

const (
	BandparametersV1610AddsrsAntennaswitchingR16Addsrs1t2rR16EnumeratedNothing = iota
	BandparametersV1610AddsrsAntennaswitchingR16Addsrs1t2rR16EnumeratedSupported
)
