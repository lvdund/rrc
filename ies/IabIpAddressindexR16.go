package ies

import "rrc/utils"

// IAB-IP-AddressIndex-r16 ::= utils.INTEGER (1..maxIAB-IP-Address-r16)
type IabIpAddressindexR16 struct {
	Value utils.INTEGER `lb:0,ub:maxIABIpAddressR16`
}
