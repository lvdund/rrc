package ies

import "rrc/utils"

// PosSystemInformation-r15-IEs ::= SEQUENCE
// Extensible
type PossysteminformationR15 struct {
	PossibTypeandinfoR15     []PossysteminformationR15IesPossibTypeandinfoR15Item `lb:1,ub:maxSIB`
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *PossysteminformationR15IesNoncriticalextension
}
