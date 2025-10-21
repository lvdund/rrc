package ies

import "rrc/utils"

// WLAN-IW-Parameters-r12 ::= SEQUENCE
type WlanIwParametersR12 struct {
	WlanIwRanRulesR12      *utils.ENUMERATED
	WlanIwAndsfPoliciesR12 *utils.ENUMERATED
}
