package ies

import "rrc/utils"

// SCGFailureInformationNR-v1590-IEs ::= SEQUENCE
type ScgfailureinformationnrV1590 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *ScgfailureinformationnrV1590IesNoncriticalextension
}
