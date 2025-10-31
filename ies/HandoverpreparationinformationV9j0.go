package ies

import "rrc/utils"

// HandoverPreparationInformation-v9j0-IEs ::= SEQUENCE
type HandoverpreparationinformationV9j0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *HandoverpreparationinformationV10j0
}
