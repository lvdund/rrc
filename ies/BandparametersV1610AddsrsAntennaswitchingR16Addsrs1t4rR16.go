package ies

import "rrc/utils"

// BandParameters-v1610-addSRS-AntennaSwitching-r16-addSRS-1T4R-r16 ::= ENUMERATED
type BandparametersV1610AddsrsAntennaswitchingR16Addsrs1t4rR16 struct {
	Value utils.ENUMERATED
}

const (
	BandparametersV1610AddsrsAntennaswitchingR16Addsrs1t4rR16EnumeratedNothing = iota
	BandparametersV1610AddsrsAntennaswitchingR16Addsrs1t4rR16EnumeratedSupported
)
