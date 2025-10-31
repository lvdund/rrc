package ies

import "rrc/utils"

// BandParameters-v1610-addSRS-AntennaSwitching-r16-addSRS-2T4R-2pairs-r16 ::= ENUMERATED
type BandparametersV1610AddsrsAntennaswitchingR16Addsrs2t4r2pairsR16 struct {
	Value utils.ENUMERATED
}

const (
	BandparametersV1610AddsrsAntennaswitchingR16Addsrs2t4r2pairsR16EnumeratedNothing = iota
	BandparametersV1610AddsrsAntennaswitchingR16Addsrs2t4r2pairsR16EnumeratedSupported
)
