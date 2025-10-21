package ies

import "rrc/utils"

// SCGFailureInformation-r12-IEs ::= SEQUENCE
type ScgfailureinformationR12Ies struct {
	FailurereportscgR12  *FailurereportscgR12
	Noncriticalextension *ScgfailureinformationV12d0aIes
}
