package ies

import "rrc/utils"

// SL-V2X-CommFreqList-r14 ::= SEQUENCE OF utils.INTEGER // SIZE (1..maxFreqV2X-r14)
type SlV2xCommfreqlistR14 struct {
	Value []utils.INTEGER `lb:1,ub:maxFreqV2XR14`
}
