package ies

import "rrc/utils"

// UECapabilityEnquiry-v1510-IEs ::= SEQUENCE
type UecapabilityenquiryV1510 struct {
	RequestedfreqbandsnrMrdcR15 *utils.OCTETSTRING
	Noncriticalextension        *UecapabilityenquiryV1530
}
