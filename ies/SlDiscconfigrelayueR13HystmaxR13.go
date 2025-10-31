package ies

import "rrc/utils"

// SL-DiscConfigRelayUE-r13-hystMax-r13 ::= ENUMERATED
type SlDiscconfigrelayueR13HystmaxR13 struct {
	Value utils.ENUMERATED
}

const (
	SlDiscconfigrelayueR13HystmaxR13EnumeratedNothing = iota
	SlDiscconfigrelayueR13HystmaxR13EnumeratedDb0
	SlDiscconfigrelayueR13HystmaxR13EnumeratedDb3
	SlDiscconfigrelayueR13HystmaxR13EnumeratedDb6
	SlDiscconfigrelayueR13HystmaxR13EnumeratedDb9
	SlDiscconfigrelayueR13HystmaxR13EnumeratedDb12
	SlDiscconfigrelayueR13HystmaxR13EnumeratedDbinf
)
