package ies

// PUCCH-ConfigDedicated-v1130 ::= SEQUENCE
type PucchConfigdedicatedV1130 struct {
	N1pucchAnCsV1130 *PucchConfigdedicatedV1130N1pucchAnCsV1130
	NpucchParamR11   *PucchConfigdedicatedV1130NpucchParamR11
}
