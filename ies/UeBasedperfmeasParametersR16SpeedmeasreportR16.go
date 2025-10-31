package ies

import "rrc/utils"

// UE-BasedPerfMeas-Parameters-r16-speedMeasReport-r16 ::= ENUMERATED
type UeBasedperfmeasParametersR16SpeedmeasreportR16 struct {
	Value utils.ENUMERATED
}

const (
	UeBasedperfmeasParametersR16SpeedmeasreportR16EnumeratedNothing = iota
	UeBasedperfmeasParametersR16SpeedmeasreportR16EnumeratedSupported
)
