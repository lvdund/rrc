package ies

import "rrc/utils"

// SIB17-IEs-r17 ::= SEQUENCE
// Extensible
type Sib17IesR17 struct {
	TrsResourcesetconfigR17  []TrsResourcesetR17 `lb:1,ub:maxNrofTRSResourcesetsR17`
	ValiditydurationR17      *Sib17IesR17ValiditydurationR17
	Latenoncriticalextension *utils.OCTETSTRING
}
