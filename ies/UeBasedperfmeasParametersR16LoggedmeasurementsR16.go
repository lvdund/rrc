package ies

import "rrc/utils"

// UE-BasedPerfMeas-Parameters-r16-loggedMeasurements-r16 ::= ENUMERATED
type UeBasedperfmeasParametersR16LoggedmeasurementsR16 struct {
	Value utils.ENUMERATED
}

const (
	UeBasedperfmeasParametersR16LoggedmeasurementsR16EnumeratedNothing = iota
	UeBasedperfmeasParametersR16LoggedmeasurementsR16EnumeratedSupported
)
