package ies

import "rrc/utils"

// CSFBParametersResponseCDMA2000-v8a0-IEs ::= SEQUENCE
type Csfbparametersresponsecdma2000V8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *Csfbparametersresponsecdma2000V8a0IesNoncriticalextension
}
