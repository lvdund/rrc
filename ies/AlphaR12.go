package ies

import "rrc/utils"

// Alpha-r12 ::= ENUMERATED
type AlphaR12 struct {
	Value utils.ENUMERATED
}

const (
	AlphaR12EnumeratedNothing = iota
	AlphaR12EnumeratedAl0
	AlphaR12EnumeratedAl04
	AlphaR12EnumeratedAl05
	AlphaR12EnumeratedAl06
	AlphaR12EnumeratedAl07
	AlphaR12EnumeratedAl08
	AlphaR12EnumeratedAl09
	AlphaR12EnumeratedAl1
)
