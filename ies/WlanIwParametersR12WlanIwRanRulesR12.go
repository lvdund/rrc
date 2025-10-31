package ies

import "rrc/utils"

// WLAN-IW-Parameters-r12-wlan-IW-RAN-Rules-r12 ::= ENUMERATED
type WlanIwParametersR12WlanIwRanRulesR12 struct {
	Value utils.ENUMERATED
}

const (
	WlanIwParametersR12WlanIwRanRulesR12EnumeratedNothing = iota
	WlanIwParametersR12WlanIwRanRulesR12EnumeratedSupported
)
