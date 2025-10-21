package ies

import "rrc/utils"

// SCGFailureInformationNR-v1590-IEs ::= SEQUENCE
type ScgfailureinformationnrV1590Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
