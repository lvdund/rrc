package ies

import "rrc/utils"

// HandoverPreparationInformation-v9d0-IEs ::= SEQUENCE
type HandoverpreparationinformationV9d0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *HandoverpreparationinformationV9e0Ies
}
