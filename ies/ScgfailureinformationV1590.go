package ies

import "rrc/utils"

// SCGFailureInformation-v1590-IEs ::= SEQUENCE
type ScgfailureinformationV1590 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *ScgfailureinformationV1590IesNoncriticalextension
}
