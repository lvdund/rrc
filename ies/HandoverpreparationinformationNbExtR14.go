package ies

import "rrc/utils"

// HandoverPreparationInformation-NB-Ext-r14-IEs ::= SEQUENCE
type HandoverpreparationinformationNbExtR14 struct {
	UeRadioaccesscapabilityinfoextR14 *utils.OCTETSTRING
	Noncriticalextension              *HandoverpreparationinformationNbExtR14IesNoncriticalextension
}
