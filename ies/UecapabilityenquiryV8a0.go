package ies

import "rrc/utils"

// UECapabilityEnquiry-v8a0-IEs ::= SEQUENCE
type UecapabilityenquiryV8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UecapabilityenquiryV1180
}
