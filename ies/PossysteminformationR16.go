package ies

import "rrc/utils"

// PosSystemInformation-r16-IEs ::= SEQUENCE
// Extensible
type PossysteminformationR16 struct {
	PossibTypeandinfoR16     []PossysteminformationR16IesPossibTypeandinfoR16Item `lb:1,ub:maxSIB`
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *PossysteminformationR16IesNoncriticalextension
}
