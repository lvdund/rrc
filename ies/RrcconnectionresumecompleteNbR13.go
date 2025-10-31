package ies

import "rrc/utils"

// RRCConnectionResumeComplete-NB-r13-IEs ::= SEQUENCE
type RrcconnectionresumecompleteNbR13 struct {
	SelectedplmnIdentityR13  *utils.INTEGER `lb:0,ub:maxPLMNR11`
	DedicatedinfonasR13      *Dedicatedinfonas
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionresumecompleteNbV1470
}
