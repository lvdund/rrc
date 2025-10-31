package ies

// BandParametersDL-r10 ::= SEQUENCE OF CA-MIMO-ParametersDL-r10
// SIZE (1..maxBandwidthClass-r10)
type BandparametersdlR10 struct {
	Value []CaMimoParametersdlR10 `lb:1,ub:maxBandwidthClassR10`
}
