package ies

import "rrc/utils"

// UERadioPagingInformation-NB-IEs ::= SEQUENCE
type UeradiopaginginformationNbIes struct {
	UeRadiopaginginfoR13 utils.OCTETSTRING
	Noncriticalextension *interface{}
}
