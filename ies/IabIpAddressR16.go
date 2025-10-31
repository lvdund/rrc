package ies

import "rrc/utils"

// IAB-IP-Address-r16 ::= CHOICE
// Extensible
const (
	IabIpAddressR16ChoiceNothing = iota
	IabIpAddressR16ChoiceIpv4AddressR16
	IabIpAddressR16ChoiceIpv6AddressR16
	IabIpAddressR16ChoiceIpv6PrefixR16
)

type IabIpAddressR16 struct {
	Choice         uint64
	Ipv4AddressR16 *utils.BITSTRING `lb:32,ub:32`
	Ipv6AddressR16 *utils.BITSTRING `lb:128,ub:128`
	Ipv6PrefixR16  *utils.BITSTRING `lb:64,ub:64`
}
