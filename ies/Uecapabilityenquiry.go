package ies

import "rrc/utils"

// UECapabilityEnquiry-IEs ::= SEQUENCE
type Uecapabilityenquiry struct {
	UeCapabilityratRequestlist UeCapabilityratRequestlist
	Latenoncriticalextension   *utils.OCTETSTRING
	UeCapabilityenquiryext     *utils.OCTETSTRING
}
