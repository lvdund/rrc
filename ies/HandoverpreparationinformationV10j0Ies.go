package ies

import "rrc/utils"

// HandoverPreparationInformation-v10j0-IEs ::= SEQUENCE
type HandoverpreparationinformationV10j0Ies struct {
	AsConfigV10j0        *AsConfigV10j0
	Noncriticalextension *HandoverpreparationinformationV10x0Ies
}
