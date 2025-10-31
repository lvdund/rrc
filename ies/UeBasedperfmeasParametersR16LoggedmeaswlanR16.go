package ies

import "rrc/utils"

// UE-BasedPerfMeas-Parameters-r16-loggedMeasWLAN-r16 ::= ENUMERATED
type UeBasedperfmeasParametersR16LoggedmeaswlanR16 struct {
	Value utils.ENUMERATED
}

const (
	UeBasedperfmeasParametersR16LoggedmeaswlanR16EnumeratedNothing = iota
	UeBasedperfmeasParametersR16LoggedmeaswlanR16EnumeratedSupported
)
