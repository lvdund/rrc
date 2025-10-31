package ies

// IAB-IP-AddressConfigurationList-r16 ::= SEQUENCE
// Extensible
type IabIpAddressconfigurationlistR16 struct {
	IabIpAddresstoaddmodlistR16  *[]IabIpAddressconfigurationR16 `lb:1,ub:maxIABIpAddressR16`
	IabIpAddresstoreleaselistR16 *[]IabIpAddressindexR16         `lb:1,ub:maxIABIpAddressR16`
}
