package ies

import "rrc/utils"

// CSFBParametersResponseCDMA2000-r8-IEs ::= SEQUENCE
type Csfbparametersresponsecdma2000R8Ies struct {
	Rand                 RandCdma2000
	Mobilityparameters   Mobilityparameterscdma2000
	Noncriticalextension *Csfbparametersresponsecdma2000V8a0Ies
}
