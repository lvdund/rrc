package ies

import "rrc/utils"

// SL-InterestedFreqList-r16 ::= SEQUENCE OF utils.INTEGER // SIZE (1..maxNrofFreqSL-r16)
type SlInterestedfreqlistR16 struct {
	Value []utils.INTEGER `lb:1,ub:maxNrofFreqSLR16`
}
