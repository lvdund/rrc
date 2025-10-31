package ies

import "rrc/utils"

// UE-BasedPerfMeas-Parameters-r16-immMeasBT-r16 ::= ENUMERATED
type UeBasedperfmeasParametersR16ImmmeasbtR16 struct {
	Value utils.ENUMERATED
}

const (
	UeBasedperfmeasParametersR16ImmmeasbtR16EnumeratedNothing = iota
	UeBasedperfmeasParametersR16ImmmeasbtR16EnumeratedSupported
)
