package ies

import "rrc/utils"

// UE-BasedNetwPerfMeasParameters-v1530-immMeasWLAN-r15 ::= ENUMERATED
type UeBasednetwperfmeasparametersV1530ImmmeaswlanR15 struct {
	Value utils.ENUMERATED
}

const (
	UeBasednetwperfmeasparametersV1530ImmmeaswlanR15EnumeratedNothing = iota
	UeBasednetwperfmeasparametersV1530ImmmeaswlanR15EnumeratedSupported
)
