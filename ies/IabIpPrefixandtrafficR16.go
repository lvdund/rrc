package ies

// IAB-IP-PrefixAndTraffic-r16 ::= SEQUENCE
type IabIpPrefixandtrafficR16 struct {
	AllTrafficIabIpAddressR16 *IabIpAddressR16
	F1CTrafficIpAddressR16    *IabIpAddressR16
	F1UTrafficIpAddressR16    *IabIpAddressR16
	NonF1TrafficIpAddressR16  *IabIpAddressR16
}
