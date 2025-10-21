package ies

import "rrc/utils"

// UECapabilityEnquiry-v1560-IEs ::= SEQUENCE
type UecapabilityenquiryV1560Ies struct {
	RequestedcapabilitycommonR15 *utils.OCTETSTRING
	Noncriticalextension         *UecapabilityenquiryV1610Ies
}
