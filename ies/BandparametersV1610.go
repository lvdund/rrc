package ies

// BandParameters-v1610 ::= SEQUENCE
type BandparametersV1610 struct {
	IntrafreqdapsR16                  *BandparametersV1610IntrafreqdapsR16
	AddsrsFrequencyhoppingR16         *BandparametersV1610AddsrsFrequencyhoppingR16
	AddsrsAntennaswitchingR16         *BandparametersV1610AddsrsAntennaswitchingR16
	SrsCapabilityperbandpairlistV1610 *[]SrsCapabilityperbandpairV1610 `lb:1,ub:maxSimultaneousBandsR10`
}
