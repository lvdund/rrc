package ies

import "rrc/utils"

// SIB4 ::= SEQUENCE
// Extensible
type Sib4 struct {
	Interfreqcarrierfreqlist      Interfreqcarrierfreqlist
	Latenoncriticalextension      *utils.OCTETSTRING
	InterfreqcarrierfreqlistV1610 *InterfreqcarrierfreqlistV1610 `ext`
	InterfreqcarrierfreqlistV1700 *InterfreqcarrierfreqlistV1700 `ext`
	InterfreqcarrierfreqlistV1720 *InterfreqcarrierfreqlistV1720 `ext`
	InterfreqcarrierfreqlistV1730 *InterfreqcarrierfreqlistV1730 `ext`
}
