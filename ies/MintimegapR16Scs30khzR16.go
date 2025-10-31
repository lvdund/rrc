package ies

import "rrc/utils"

// MinTimeGap-r16-scs-30kHz-r16 ::= ENUMERATED
type MintimegapR16Scs30khzR16 struct {
	Value utils.ENUMERATED
}

const (
	MintimegapR16Scs30khzR16EnumeratedNothing = iota
	MintimegapR16Scs30khzR16EnumeratedSl1
	MintimegapR16Scs30khzR16EnumeratedSl6
)
