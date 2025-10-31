package ies

import "rrc/utils"

// WUS-Config-NB-r15-numPOs-r15 ::= ENUMERATED
type WusConfigNbR15NumposR15 struct {
	Value utils.ENUMERATED
}

const (
	WusConfigNbR15NumposR15EnumeratedNothing = iota
	WusConfigNbR15NumposR15EnumeratedN1
	WusConfigNbR15NumposR15EnumeratedN2
	WusConfigNbR15NumposR15EnumeratedN4
)
