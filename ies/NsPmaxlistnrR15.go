package ies

import "rrc/utils"

// NS-PmaxListNR-r15 ::= SEQUENCE OF NS-PmaxValueNR-r15
// SIZE (1..8)
type NsPmaxlistnrR15 struct {
	Value []NsPmaxvaluenrR15
}
