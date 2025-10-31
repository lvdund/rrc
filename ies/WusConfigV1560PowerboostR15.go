package ies

import "rrc/utils"

// WUS-Config-v1560-powerBoost-r15 ::= ENUMERATED
type WusConfigV1560PowerboostR15 struct {
	Value utils.ENUMERATED
}

const (
	WusConfigV1560PowerboostR15EnumeratedNothing = iota
	WusConfigV1560PowerboostR15EnumeratedDb0
	WusConfigV1560PowerboostR15EnumeratedDb1dot8
	WusConfigV1560PowerboostR15EnumeratedDb3
	WusConfigV1560PowerboostR15EnumeratedDb4dot8
)
