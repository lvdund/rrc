package ies

// NPRACH-ConfigSIB-NB-r13 ::= SEQUENCE
type NprachConfigsibNbR13 struct {
	NprachCpLengthR13              NprachConfigsibNbR13NprachCpLengthR13
	RsrpThresholdsprachinfolistR13 *RsrpThresholdsnprachInfolistNbR13
	NprachParameterslistR13        NprachParameterslistNbR13
}
