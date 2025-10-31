package ies

import "rrc/utils"

// SCGFailureInformation-v12d0a-IEs ::= SEQUENCE
type ScgfailureinformationV12d0a struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *ScgfailureinformationV12d0aIesNoncriticalextension
}
