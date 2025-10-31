package ies

import "rrc/utils"

// MinTimeGap-r16-scs-15kHz-r16 ::= ENUMERATED
type MintimegapR16Scs15khzR16 struct {
	Value utils.ENUMERATED
}

const (
	MintimegapR16Scs15khzR16EnumeratedNothing = iota
	MintimegapR16Scs15khzR16EnumeratedSl1
	MintimegapR16Scs15khzR16EnumeratedSl3
)
