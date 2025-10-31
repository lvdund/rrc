package ies

import "rrc/utils"

// SystemInformation-NB-r13-IEs ::= SEQUENCE
// Extensible
type SysteminformationNbR13 struct {
	SibTypeandinfoR13        []SysteminformationNbR13IesSibTypeandinfoR13Item `lb:1,ub:maxSIB`
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *SysteminformationNbR13IesNoncriticalextension
}
