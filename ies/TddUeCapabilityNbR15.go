package ies

import "rrc/utils"

// TDD-UE-Capability-NB-r15 ::= SEQUENCE
// Extensible
type TddUeCapabilityNbR15 struct {
	UeCategoryNbR15            *utils.ENUMERATED
	Phylayerparametersrel13R15 *PhylayerparametersNbR13
	Phylayerparametersrel14R15 *PhylayerparametersNbV1430
	PhylayerparametersV1530    *PhylayerparametersNbV1530
}
