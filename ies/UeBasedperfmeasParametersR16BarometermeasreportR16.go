package ies

import "rrc/utils"

// UE-BasedPerfMeas-Parameters-r16-barometerMeasReport-r16 ::= ENUMERATED
type UeBasedperfmeasParametersR16BarometermeasreportR16 struct {
	Value utils.ENUMERATED
}

const (
	UeBasedperfmeasParametersR16BarometermeasreportR16EnumeratedNothing = iota
	UeBasedperfmeasParametersR16BarometermeasreportR16EnumeratedSupported
)
