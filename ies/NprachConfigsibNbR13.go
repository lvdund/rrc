package ies

import "rrc/utils"

// NPRACH-ConfigSIB-NB-r13 ::= SEQUENCE
type NprachConfigsibNbR13 struct {
	NprachCpLengthR13              utils.ENUMERATED
	RsrpThresholdsprachinfolistR13 *RsrpThresholdsnprachInfolistNbR13
	NprachParameterslistR13        NprachParameterslistNbR13
}
