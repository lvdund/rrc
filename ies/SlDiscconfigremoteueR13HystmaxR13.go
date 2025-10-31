package ies

import "rrc/utils"

// SL-DiscConfigRemoteUE-r13-hystMax-r13 ::= ENUMERATED
type SlDiscconfigremoteueR13HystmaxR13 struct {
	Value utils.ENUMERATED
}

const (
	SlDiscconfigremoteueR13HystmaxR13EnumeratedNothing = iota
	SlDiscconfigremoteueR13HystmaxR13EnumeratedDb0
	SlDiscconfigremoteueR13HystmaxR13EnumeratedDb3
	SlDiscconfigremoteueR13HystmaxR13EnumeratedDb6
	SlDiscconfigremoteueR13HystmaxR13EnumeratedDb9
	SlDiscconfigremoteueR13HystmaxR13EnumeratedDb12
)
