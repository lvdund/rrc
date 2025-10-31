package ies

import "rrc/utils"

// CA-ParametersNR-v1610-crossCarrierSchedulingDefaultQCL-r16 ::= ENUMERATED
type CaParametersnrV1610CrosscarrierschedulingdefaultqclR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610CrosscarrierschedulingdefaultqclR16EnumeratedNothing = iota
	CaParametersnrV1610CrosscarrierschedulingdefaultqclR16EnumeratedDiff_Only
	CaParametersnrV1610CrosscarrierschedulingdefaultqclR16EnumeratedBoth
)
