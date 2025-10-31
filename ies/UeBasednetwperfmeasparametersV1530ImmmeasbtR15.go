package ies

import "rrc/utils"

// UE-BasedNetwPerfMeasParameters-v1530-immMeasBT-r15 ::= ENUMERATED
type UeBasednetwperfmeasparametersV1530ImmmeasbtR15 struct {
	Value utils.ENUMERATED
}

const (
	UeBasednetwperfmeasparametersV1530ImmmeasbtR15EnumeratedNothing = iota
	UeBasednetwperfmeasparametersV1530ImmmeasbtR15EnumeratedSupported
)
