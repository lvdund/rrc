package ies

import "rrc/utils"

// WLAN-Status-v1430 ::= ENUMERATED
type WlanStatusV1430 struct {
	Value utils.ENUMERATED
}

const (
	WlanStatusV1430EnumeratedNothing = iota
	WlanStatusV1430EnumeratedSuspended
	WlanStatusV1430EnumeratedResumed
)
