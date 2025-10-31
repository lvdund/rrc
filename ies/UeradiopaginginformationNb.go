package ies

import "rrc/utils"

// UERadioPagingInformation-NB-IEs ::= SEQUENCE
type UeradiopaginginformationNb struct {
	UeRadiopaginginfoR13 utils.OCTETSTRING
	Noncriticalextension *UeradiopaginginformationNbIesNoncriticalextension
}
