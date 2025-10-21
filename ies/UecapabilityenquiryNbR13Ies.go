package ies

import "rrc/utils"

// UECapabilityEnquiry-NB-r13-IEs ::= SEQUENCE
type UecapabilityenquiryNbR13Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
