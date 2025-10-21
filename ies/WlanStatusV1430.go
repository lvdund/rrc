package ies

import "rrc/utils"

// WLAN-Status-v1430 ::= ENUMERATED
type WlanStatusV1430 struct {
	Value utils.ENUMERATED
}

const (
	WlanStatusV1430Suspended = 0
	WlanStatusV1430Resumed   = 1
)
