package ies

// BandParameters-v1270 ::= SEQUENCE
type BandparametersV1270 struct {
	BandparametersdlV1270 []CaMimoParametersdlV1270 `lb:1,ub:maxBandwidthClassR10`
}
