package ies

import "rrc/utils"

// UECapabilityEnquiry-v1180-IEs ::= SEQUENCE
type UecapabilityenquiryV1180Ies struct {
	RequestedfrequencybandsR11 *interface{}
	Noncriticalextension       *UecapabilityenquiryV1310Ies
}
