package ies

import "rrc/utils"

// UECapabilityEnquiry-v8a0-IEs ::= SEQUENCE
type UecapabilityenquiryV8a0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UecapabilityenquiryV1180Ies
}
