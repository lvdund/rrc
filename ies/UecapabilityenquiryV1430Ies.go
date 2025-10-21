package ies

import "rrc/utils"

// UECapabilityEnquiry-v1430-IEs ::= SEQUENCE
type UecapabilityenquiryV1430Ies struct {
	RequestdifffallbackcomblistR14 *BandcombinationlistR14
	Noncriticalextension           *UecapabilityenquiryV1510Ies
}
