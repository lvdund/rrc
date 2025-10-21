package ies

import "rrc/utils"

// WLAN-Status-r13 ::= ENUMERATED
type WlanStatusR13 struct {
	Value utils.ENUMERATED
}

const (
	WlanStatusR13Successfulassociation  = 0
	WlanStatusR13Failurewlanradiolink   = 1
	WlanStatusR13Failurewlanunavailable = 2
	WlanStatusR13Failuretimeout         = 3
)
