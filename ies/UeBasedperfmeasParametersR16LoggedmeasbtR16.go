package ies

import "rrc/utils"

// UE-BasedPerfMeas-Parameters-r16-loggedMeasBT-r16 ::= ENUMERATED
type UeBasedperfmeasParametersR16LoggedmeasbtR16 struct {
	Value utils.ENUMERATED
}

const (
	UeBasedperfmeasParametersR16LoggedmeasbtR16EnumeratedNothing = iota
	UeBasedperfmeasParametersR16LoggedmeasbtR16EnumeratedSupported
)
