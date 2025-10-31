package ies

// TDM-AssistanceInfo-r11 ::= CHOICE
// Extensible
const (
	TdmAssistanceinfoR11ChoiceNothing = iota
	TdmAssistanceinfoR11ChoiceDrxAssistanceinfoR11
	TdmAssistanceinfoR11ChoiceIdcSubframepatternlistR11
)

type TdmAssistanceinfoR11 struct {
	Choice                    uint64
	DrxAssistanceinfoR11      *TdmAssistanceinfoR11DrxAssistanceinfoR11
	IdcSubframepatternlistR11 *IdcSubframepatternlistR11
}
