package ies

import "rrc/utils"

// TDM-AssistanceInfo-r11-drx-AssistanceInfo-r11 ::= SEQUENCE
type TdmAssistanceinfoR11DrxAssistanceinfoR11 struct {
	DrxCyclelengthR11 TdmAssistanceinfoR11DrxAssistanceinfoR11DrxCyclelengthR11
	DrxOffsetR11      *utils.INTEGER `lb:0,ub:255`
	DrxActivetimeR11  TdmAssistanceinfoR11DrxAssistanceinfoR11DrxActivetimeR11
}
