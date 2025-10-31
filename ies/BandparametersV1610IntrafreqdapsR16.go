package ies

// BandParameters-v1610-intraFreqDAPS-r16 ::= SEQUENCE
type BandparametersV1610IntrafreqdapsR16 struct {
	IntrafreqasyncdapsR16   *BandparametersV1610IntrafreqdapsR16IntrafreqasyncdapsR16
	Dummy                   *BandparametersV1610IntrafreqdapsR16Dummy
	IntrafreqtwotagsDapsR16 *BandparametersV1610IntrafreqdapsR16IntrafreqtwotagsDapsR16
}
