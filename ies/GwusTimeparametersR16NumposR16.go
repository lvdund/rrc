package ies

import "rrc/utils"

// GWUS-TimeParameters-r16-numPOs-r16 ::= ENUMERATED
type GwusTimeparametersR16NumposR16 struct {
	Value utils.ENUMERATED
}

const (
	GwusTimeparametersR16NumposR16EnumeratedNothing = iota
	GwusTimeparametersR16NumposR16EnumeratedN1
	GwusTimeparametersR16NumposR16EnumeratedN2
	GwusTimeparametersR16NumposR16EnumeratedN4
	GwusTimeparametersR16NumposR16EnumeratedSpare1
)
