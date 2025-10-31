package ies

// BandParameters-v10i0 ::= SEQUENCE
type BandparametersV10i0 struct {
	BandparametersdlV10i0 []CaMimoParametersdlV10i0 `lb:1,ub:maxBandwidthClassR10`
}
