package ies

import "rrc/utils"

// MinTimeGap-r16-scs-120kHz-r16 ::= ENUMERATED
type MintimegapR16Scs120khzR16 struct {
	Value utils.ENUMERATED
}

const (
	MintimegapR16Scs120khzR16EnumeratedNothing = iota
	MintimegapR16Scs120khzR16EnumeratedSl2
	MintimegapR16Scs120khzR16EnumeratedSl24
)
