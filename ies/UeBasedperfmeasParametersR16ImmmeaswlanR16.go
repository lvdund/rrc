package ies

import "rrc/utils"

// UE-BasedPerfMeas-Parameters-r16-immMeasWLAN-r16 ::= ENUMERATED
type UeBasedperfmeasParametersR16ImmmeaswlanR16 struct {
	Value utils.ENUMERATED
}

const (
	UeBasedperfmeasParametersR16ImmmeaswlanR16EnumeratedNothing = iota
	UeBasedperfmeasParametersR16ImmmeaswlanR16EnumeratedSupported
)
