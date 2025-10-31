package ies

import "rrc/utils"

// WLAN-Status-r13 ::= ENUMERATED
type WlanStatusR13 struct {
	Value utils.ENUMERATED
}

const (
	WlanStatusR13EnumeratedNothing = iota
	WlanStatusR13EnumeratedSuccessfulassociation
	WlanStatusR13EnumeratedFailurewlanradiolink
	WlanStatusR13EnumeratedFailurewlanunavailable
	WlanStatusR13EnumeratedFailuretimeout
)
