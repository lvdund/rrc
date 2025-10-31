package ies

import "rrc/utils"

// SIB20-r17 ::= SEQUENCE
// Extensible
type Sib20R17 struct {
	McchConfigR17            McchConfigR17
	CfrConfigmcchMtchR17     *CfrConfigmcchMtchR17
	Latenoncriticalextension *utils.OCTETSTRING
}
