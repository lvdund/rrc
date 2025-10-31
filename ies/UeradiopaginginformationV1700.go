package ies

import "rrc/utils"

// UERadioPagingInformation-v1700-IEs ::= SEQUENCE
type UeradiopaginginformationV1700 struct {
	UeRadiopaginginfoR17            *utils.OCTETSTRING
	InactivestatepoDeterminationR17 *UeradiopaginginformationV1700IesInactivestatepoDeterminationR17
	NumberofrxredcapR17             *UeradiopaginginformationV1700IesNumberofrxredcapR17
	HalfduplexfddTypeaRedcapR17     *[]Freqbandindicatornr `lb:1,ub:maxBands`
	Noncriticalextension            *UeradiopaginginformationV1700IesNoncriticalextension
}
