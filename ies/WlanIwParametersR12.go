package ies

// WLAN-IW-Parameters-r12 ::= SEQUENCE
type WlanIwParametersR12 struct {
	WlanIwRanRulesR12      *WlanIwParametersR12WlanIwRanRulesR12
	WlanIwAndsfPoliciesR12 *WlanIwParametersR12WlanIwAndsfPoliciesR12
}
