package ies

import "rrc/utils"

// SystemInformationBlockType29-r16 ::= SEQUENCE
// Extensible
type Systeminformationblocktype29R16 struct {
	ResourcereservationconfigcommondlR16 *ResourcereservationconfigdlR16
	ResourcereservationconfigcommonulR16 *ResourcereservationconfigulR16
	Latenoncriticalextension             *utils.OCTETSTRING
}
