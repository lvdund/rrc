package ies

import "rrc/utils"

// FailureInformation-IEs ::= SEQUENCE
type Failureinformation struct {
	FailureinforlcBearer     *FailureinforlcBearer
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *FailureinformationV1610
}
