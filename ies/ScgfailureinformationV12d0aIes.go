package ies

import "rrc/utils"

// SCGFailureInformation-v12d0a-IEs ::= SEQUENCE
type ScgfailureinformationV12d0aIes struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
