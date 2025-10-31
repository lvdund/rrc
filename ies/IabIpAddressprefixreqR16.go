package ies

import "rrc/utils"

// IAB-IP-AddressPrefixReq-r16 ::= SEQUENCE
// Extensible
type IabIpAddressprefixreqR16 struct {
	AllTrafficPrefixreqR16   *utils.BOOLEAN
	F1CTrafficPrefixreqR16   *utils.BOOLEAN
	F1UTrafficPrefixreqR16   *utils.BOOLEAN
	NonF1TrafficPrefixreqR16 *utils.BOOLEAN
}
