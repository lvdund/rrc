package ies

import "rrc/utils"

// IP-Address-r13 ::= CHOICE
const (
	IpAddressR13ChoiceNothing = iota
	IpAddressR13ChoiceIpv4R13
	IpAddressR13ChoiceIpv6R13
)

type IpAddressR13 struct {
	Choice  uint64
	Ipv4R13 *utils.BITSTRING `lb:32,ub:32`
	Ipv6R13 *utils.BITSTRING `lb:128,ub:128`
}
