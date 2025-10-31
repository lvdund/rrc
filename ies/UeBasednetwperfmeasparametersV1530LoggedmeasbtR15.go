package ies

import "rrc/utils"

// UE-BasedNetwPerfMeasParameters-v1530-loggedMeasBT-r15 ::= ENUMERATED
type UeBasednetwperfmeasparametersV1530LoggedmeasbtR15 struct {
	Value utils.ENUMERATED
}

const (
	UeBasednetwperfmeasparametersV1530LoggedmeasbtR15EnumeratedNothing = iota
	UeBasednetwperfmeasparametersV1530LoggedmeasbtR15EnumeratedSupported
)
