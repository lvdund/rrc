package ies

import "rrc/utils"

// IP-Address-r13 ::= CHOICE
type IpAddressR13 interface {
	isIpAddressR13()
}

type IpAddressR13Ipv4R13 struct {
	Value utils.BITSTRING
}

func (*IpAddressR13Ipv4R13) isIpAddressR13() {}

type IpAddressR13Ipv6R13 struct {
	Value utils.BITSTRING
}

func (*IpAddressR13Ipv6R13) isIpAddressR13() {}
