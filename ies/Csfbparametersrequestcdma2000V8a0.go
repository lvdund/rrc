package ies

import "rrc/utils"

// CSFBParametersRequestCDMA2000-v8a0-IEs ::= SEQUENCE
type Csfbparametersrequestcdma2000V8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *Csfbparametersrequestcdma2000V8a0IesNoncriticalextension
}
