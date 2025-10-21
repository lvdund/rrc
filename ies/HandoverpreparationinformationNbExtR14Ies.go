package ies

import "rrc/utils"

// HandoverPreparationInformation-NB-Ext-r14-IEs ::= SEQUENCE
type HandoverpreparationinformationNbExtR14Ies struct {
	UeRadioaccesscapabilityinfoextR14 *utils.OCTETSTRING
	Noncriticalextension              *interface{}
}
