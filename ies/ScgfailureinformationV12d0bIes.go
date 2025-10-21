package ies

import "rrc/utils"

// SCGFailureInformation-v12d0b-IEs ::= SEQUENCE
type ScgfailureinformationV12d0bIes struct {
	FailurereportscgV12d0 *FailurereportscgV12d0
	Noncriticalextension  *interface{}
}
