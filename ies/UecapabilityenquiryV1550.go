package ies

import "rrc/utils"

// UECapabilityEnquiry-v1550-IEs ::= SEQUENCE
type UecapabilityenquiryV1550 struct {
	RequestedcapabilitynrR15 *utils.OCTETSTRING
	Noncriticalextension     *UecapabilityenquiryV1560
}
