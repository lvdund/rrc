package ies

import "rrc/utils"

// SystemInformation-IEs ::= SEQUENCE
// Extensible
type Systeminformation struct {
	SibTypeandinfo           []SysteminformationIesSibTypeandinfoItem `lb:1,ub:maxSIB`
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *SysteminformationIesNoncriticalextension
}
