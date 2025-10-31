package ies

import "rrc/utils"

// HandoverPreparationInformation-v1540-IEs ::= SEQUENCE
type HandoverpreparationinformationV1540 struct {
	SourcerbConfigintra5gcR15 *utils.OCTETSTRING
	Noncriticalextension      *HandoverpreparationinformationV1610
}
