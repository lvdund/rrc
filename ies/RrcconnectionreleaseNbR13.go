package ies

import "rrc/utils"

// RRCConnectionRelease-NB-r13-IEs ::= SEQUENCE
type RrcconnectionreleaseNbR13 struct {
	ReleasecauseR13          ReleasecauseNbR13
	ResumeidentityR13        *ResumeidentityR13
	ExtendedwaittimeR13      *utils.INTEGER `lb:0,ub:1800`
	RedirectedcarrierinfoR13 *RedirectedcarrierinfoNbR13
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionreleaseNbV1430
}
