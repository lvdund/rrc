package ies

import "rrc/utils"

// MCGFailureInformation-r16-IEs ::= SEQUENCE
type McgfailureinformationR16 struct {
	FailurereportmcgR16      *FailurereportmcgR16
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *McgfailureinformationR16IesNoncriticalextension
}
