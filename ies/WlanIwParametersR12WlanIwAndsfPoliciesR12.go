package ies

import "rrc/utils"

// WLAN-IW-Parameters-r12-wlan-IW-ANDSF-Policies-r12 ::= ENUMERATED
type WlanIwParametersR12WlanIwAndsfPoliciesR12 struct {
	Value utils.ENUMERATED
}

const (
	WlanIwParametersR12WlanIwAndsfPoliciesR12EnumeratedNothing = iota
	WlanIwParametersR12WlanIwAndsfPoliciesR12EnumeratedSupported
)
