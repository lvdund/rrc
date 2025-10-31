package ies

import "rrc/utils"

// IDC-SubframePattern-r11 ::= CHOICE
// Extensible
const (
	IdcSubframepatternR11ChoiceNothing = iota
	IdcSubframepatternR11ChoiceSubframepatternfddR11
	IdcSubframepatternR11ChoiceSubframepatterntddR11
)

type IdcSubframepatternR11 struct {
	Choice                uint64
	SubframepatternfddR11 *utils.BITSTRING `lb:4,ub:4`
	SubframepatterntddR11 *IdcSubframepatternR11SubframepatterntddR11
}
