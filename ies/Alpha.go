package ies

import "rrc/utils"

// Alpha ::= ENUMERATED
type Alpha struct {
	Value utils.ENUMERATED
}

const (
	AlphaEnumeratedNothing = iota
	AlphaEnumeratedAlpha0
	AlphaEnumeratedAlpha04
	AlphaEnumeratedAlpha05
	AlphaEnumeratedAlpha06
	AlphaEnumeratedAlpha07
	AlphaEnumeratedAlpha08
	AlphaEnumeratedAlpha09
	AlphaEnumeratedAlpha1
)
