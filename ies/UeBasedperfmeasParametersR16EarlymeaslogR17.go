package ies

import "rrc/utils"

// UE-BasedPerfMeas-Parameters-r16-earlyMeasLog-r17 ::= ENUMERATED
type UeBasedperfmeasParametersR16EarlymeaslogR17 struct {
	Value utils.ENUMERATED
}

const (
	UeBasedperfmeasParametersR16EarlymeaslogR17EnumeratedNothing = iota
	UeBasedperfmeasParametersR16EarlymeaslogR17EnumeratedSupported
)
