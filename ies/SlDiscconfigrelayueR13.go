package ies

import "rrc/utils"

// SL-DiscConfigRelayUE-r13 ::= SEQUENCE
type SlDiscconfigrelayueR13 struct {
	ThreshhighR13 *RsrpRangesl4R13
	ThreshlowR13  *RsrpRangesl4R13
	HystmaxR13    *utils.ENUMERATED
	HystminR13    *utils.ENUMERATED
}
