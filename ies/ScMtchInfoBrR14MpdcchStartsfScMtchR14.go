package ies

import "rrc/utils"

// SC-MTCH-Info-BR-r14-mpdcch-StartSF-SC-MTCH-r14 ::= CHOICE
const (
	ScMtchInfoBrR14MpdcchStartsfScMtchR14ChoiceNothing = iota
	ScMtchInfoBrR14MpdcchStartsfScMtchR14ChoiceFddR14
	ScMtchInfoBrR14MpdcchStartsfScMtchR14ChoiceTddR14
)

type ScMtchInfoBrR14MpdcchStartsfScMtchR14 struct {
	Choice uint64
	FddR14 *utils.ENUMERATED
	TddR14 *utils.ENUMERATED
}
