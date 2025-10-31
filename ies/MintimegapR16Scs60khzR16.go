package ies

import "rrc/utils"

// MinTimeGap-r16-scs-60kHz-r16 ::= ENUMERATED
type MintimegapR16Scs60khzR16 struct {
	Value utils.ENUMERATED
}

const (
	MintimegapR16Scs60khzR16EnumeratedNothing = iota
	MintimegapR16Scs60khzR16EnumeratedSl1
	MintimegapR16Scs60khzR16EnumeratedSl12
)
