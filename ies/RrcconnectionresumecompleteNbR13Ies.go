package ies

import "rrc/utils"

// RRCConnectionResumeComplete-NB-r13-IEs ::= SEQUENCE
type RrcconnectionresumecompleteNbR13Ies struct {
	SelectedplmnIdentityR13  *utils.INTEGER
	DedicatedinfonasR13      *Dedicatedinfonas
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionresumecompleteNbV1470Ies
}
