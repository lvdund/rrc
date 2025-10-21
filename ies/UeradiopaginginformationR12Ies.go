package ies

import "rrc/utils"

// UERadioPagingInformation-r12-IEs ::= SEQUENCE
type UeradiopaginginformationR12Ies struct {
	UeRadiopaginginfoR12 utils.OCTETSTRING
	Noncriticalextension *UeradiopaginginformationV1310Ies
}
