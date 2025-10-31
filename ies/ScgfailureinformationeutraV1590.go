package ies

import "rrc/utils"

// SCGFailureInformationEUTRA-v1590-IEs ::= SEQUENCE
type ScgfailureinformationeutraV1590 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *ScgfailureinformationeutraV1590IesNoncriticalextension
}
