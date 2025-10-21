package ies

import "rrc/utils"

// HandoverPreparationInformation-v1430-IEs ::= SEQUENCE
type HandoverpreparationinformationV1430Ies struct {
	AsConfigV1430         *AsConfigV1430
	MakebeforebreakreqR14 *utils.ENUMERATED
	Noncriticalextension  *HandoverpreparationinformationV1530Ies
}
