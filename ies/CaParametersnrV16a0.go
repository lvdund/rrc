package ies

// CA-ParametersNR-v16a0 ::= SEQUENCE
type CaParametersnrV16a0 struct {
	PdcchBlinddetectionmixedlistR16 []PdcchBlinddetectionmixedlistR16 `lb:1,ub:maxNrofPdcchBlinddetectionmixed1R16`
}
