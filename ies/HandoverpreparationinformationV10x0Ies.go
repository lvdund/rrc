package ies

import "rrc/utils"

// HandoverPreparationInformation-v10x0-IEs ::= SEQUENCE
type HandoverpreparationinformationV10x0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *HandoverpreparationinformationV13c0Ies
}
