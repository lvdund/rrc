package ies

// BandParametersUL-r10 ::= SEQUENCE OF CA-MIMO-ParametersUL-r10
// SIZE (1..maxBandwidthClass-r10)
type BandparametersulR10 struct {
	Value []CaMimoParametersulR10 `lb:1,ub:maxBandwidthClassR10`
}
