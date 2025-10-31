package ies

import "rrc/utils"

// IAB-IP-AddressNumReq-r16 ::= SEQUENCE
// Extensible
type IabIpAddressnumreqR16 struct {
	AllTrafficNumreqR16   *utils.INTEGER `lb:0,ub:8`
	F1CTrafficNumreqR16   *utils.INTEGER `lb:0,ub:8`
	F1UTrafficNumreqR16   *utils.INTEGER `lb:0,ub:8`
	NonF1TrafficNumreqR16 *utils.INTEGER `lb:0,ub:8`
}
