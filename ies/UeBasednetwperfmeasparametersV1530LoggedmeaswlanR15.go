package ies

import "rrc/utils"

// UE-BasedNetwPerfMeasParameters-v1530-loggedMeasWLAN-r15 ::= ENUMERATED
type UeBasednetwperfmeasparametersV1530LoggedmeaswlanR15 struct {
	Value utils.ENUMERATED
}

const (
	UeBasednetwperfmeasparametersV1530LoggedmeaswlanR15EnumeratedNothing = iota
	UeBasednetwperfmeasparametersV1530LoggedmeaswlanR15EnumeratedSupported
)
