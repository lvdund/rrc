package ies

import "rrc/utils"

// IDC-SubframePattern-r11-subframePatternTDD-r11 ::= CHOICE
const (
	IdcSubframepatternR11SubframepatterntddR11ChoiceNothing = iota
	IdcSubframepatternR11SubframepatterntddR11ChoiceSubframeconfig0R11
	IdcSubframepatternR11SubframepatterntddR11ChoiceSubframeconfig15R11
	IdcSubframepatternR11SubframepatterntddR11ChoiceSubframeconfig6R11
)

type IdcSubframepatternR11SubframepatterntddR11 struct {
	Choice              uint64
	Subframeconfig0R11  *utils.BITSTRING `lb:70,ub:70`
	Subframeconfig15R11 *utils.BITSTRING `lb:10,ub:10`
	Subframeconfig6R11  *utils.BITSTRING `lb:60,ub:60`
}
