package ies

import "rrc/utils"

// IAB-IP-AddressConfiguration-r16 ::= SEQUENCE
// Extensible
type IabIpAddressconfigurationR16 struct {
	IabIpAddressindexR16    IabIpAddressindexR16
	IabIpAddressR16         *IabIpAddressR16
	IabIpUsageR16           *IabIpUsageR16
	IabDonorDuBapAddressR16 *utils.BITSTRING `lb:10,ub:10`
}
