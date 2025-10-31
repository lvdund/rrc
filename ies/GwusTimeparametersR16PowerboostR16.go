package ies

import "rrc/utils"

// GWUS-TimeParameters-r16-powerBoost-r16 ::= ENUMERATED
type GwusTimeparametersR16PowerboostR16 struct {
	Value utils.ENUMERATED
}

const (
	GwusTimeparametersR16PowerboostR16EnumeratedNothing = iota
	GwusTimeparametersR16PowerboostR16EnumeratedDb0
	GwusTimeparametersR16PowerboostR16EnumeratedDb1dot8
	GwusTimeparametersR16PowerboostR16EnumeratedDb3
	GwusTimeparametersR16PowerboostR16EnumeratedDb4dot8
)
