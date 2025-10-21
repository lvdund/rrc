package ies

import "rrc/utils"

// UE-BasedNetwPerfMeasParameters-v1530 ::= SEQUENCE
type UeBasednetwperfmeasparametersV1530 struct {
	LoggedmeasbtR15   *utils.ENUMERATED
	LoggedmeaswlanR15 *utils.ENUMERATED
	ImmmeasbtR15      *utils.ENUMERATED
	ImmmeaswlanR15    *utils.ENUMERATED
}
