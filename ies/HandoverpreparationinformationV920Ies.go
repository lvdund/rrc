package ies

import "rrc/utils"

// HandoverPreparationInformation-v920-IEs ::= SEQUENCE
// Extensible
type HandoverpreparationinformationV920Ies struct {
	UeConfigreleaseR9    *utils.ENUMERATED
	Noncriticalextension *HandoverpreparationinformationV9d0Ies
}
