package ies

// CA-MIMO-ParametersDL-v1270 ::= SEQUENCE
type CaMimoParametersdlV1270 struct {
	IntrabandcontiguousccInfolistR12 []IntrabandcontiguousccInfoR12 `lb:1,ub:maxServCellR10`
}
