package ies

import "rrc/utils"

// HandoverPreparationInformation-NB-v1380-IEs ::= SEQUENCE
type HandoverpreparationinformationNbV1380 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *HandoverpreparationinformationNbExtR14
}
