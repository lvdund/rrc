package ies

// IAB-IP-AddressAndTraffic-r16 ::= SEQUENCE
type IabIpAddressandtrafficR16 struct {
	AllTrafficIabIpAddressR16 *[]IabIpAddressR16 `lb:1,ub:8`
	F1CTrafficIpAddressR16    *[]IabIpAddressR16 `lb:1,ub:8`
	F1UTrafficIpAddressR16    *[]IabIpAddressR16 `lb:1,ub:8`
	NonF1TrafficIpAddressR16  *[]IabIpAddressR16 `lb:1,ub:8`
}
