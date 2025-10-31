package ies

import "rrc/utils"

// SL-DiscConfigRelayUE-r13-hystMin-r13 ::= ENUMERATED
type SlDiscconfigrelayueR13HystminR13 struct {
	Value utils.ENUMERATED
}

const (
	SlDiscconfigrelayueR13HystminR13EnumeratedNothing = iota
	SlDiscconfigrelayueR13HystminR13EnumeratedDb0
	SlDiscconfigrelayueR13HystminR13EnumeratedDb3
	SlDiscconfigrelayueR13HystminR13EnumeratedDb6
	SlDiscconfigrelayueR13HystminR13EnumeratedDb9
	SlDiscconfigrelayueR13HystminR13EnumeratedDb12
)
