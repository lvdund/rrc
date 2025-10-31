package ies

import "rrc/utils"

// SIB18-r17 ::= SEQUENCE
// Extensible
type Sib18R17 struct {
	GinElementlistR17        *[]GinElementR17  `lb:1,ub:maxGINR17`
	GinsPersnpnListR17       *[]GinsPersnpnR17 `lb:1,ub:maxNPNR16`
	Latenoncriticalextension *utils.OCTETSTRING
}
