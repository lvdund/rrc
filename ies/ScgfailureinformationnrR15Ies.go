package ies

import "rrc/utils"

// SCGFailureInformationNR-r15-IEs ::= SEQUENCE
type ScgfailureinformationnrR15Ies struct {
	FailurereportscgNrR15 *FailurereportscgNrR15
	Noncriticalextension  *ScgfailureinformationnrV1590Ies
}
