package ies

import "rrc/utils"

// WUS-Config-r15-numPOs-r15 ::= ENUMERATED
type WusConfigR15NumposR15 struct {
	Value utils.ENUMERATED
}

const (
	WusConfigR15NumposR15EnumeratedNothing = iota
	WusConfigR15NumposR15EnumeratedN1
	WusConfigR15NumposR15EnumeratedN2
	WusConfigR15NumposR15EnumeratedN4
	WusConfigR15NumposR15EnumeratedSpare1
)
