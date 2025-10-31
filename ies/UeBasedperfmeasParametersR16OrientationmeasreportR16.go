package ies

import "rrc/utils"

// UE-BasedPerfMeas-Parameters-r16-orientationMeasReport-r16 ::= ENUMERATED
type UeBasedperfmeasParametersR16OrientationmeasreportR16 struct {
	Value utils.ENUMERATED
}

const (
	UeBasedperfmeasParametersR16OrientationmeasreportR16EnumeratedNothing = iota
	UeBasedperfmeasParametersR16OrientationmeasreportR16EnumeratedSupported
)
