package ies

import "rrc/utils"

// MultiBandNsPmaxListNR-1-v1550 ::= SEQUENCE OF NS-PmaxListNR-r15
// SIZE (1.. maxMultiBandsNR-1-r15)
type Multibandnspmaxlistnr1V1550 struct {
	Value utils.Sequence[NsPmaxlistnrR15]
}
