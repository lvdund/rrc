package ies

import "rrc/utils"

// NS-PmaxList-NB-r13 ::= SEQUENCE OF NS-PmaxValue-NB-r13
// SIZE (1..maxNS-Pmax-NB-r13)
type NsPmaxlistNbR13 struct {
	Value utils.Sequence[NsPmaxvalueNbR13]
}
