package ies

import "rrc/utils"

// UECapabilityEnquiry-NB-r13-IEs ::= SEQUENCE
type UecapabilityenquiryNbR13 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UecapabilityenquiryNbR13IesNoncriticalextension
}
