package ies

import "rrc/utils"

// DRX-Config-v1130-longDRX-CycleStartOffset-v1130 ::= CHOICE
const (
	DrxConfigV1130LongdrxCyclestartoffsetV1130ChoiceNothing = iota
	DrxConfigV1130LongdrxCyclestartoffsetV1130ChoiceSf60V1130
	DrxConfigV1130LongdrxCyclestartoffsetV1130ChoiceSf70V1130
)

type DrxConfigV1130LongdrxCyclestartoffsetV1130 struct {
	Choice    uint64
	Sf60V1130 *utils.INTEGER `lb:0,ub:59`
	Sf70V1130 *utils.INTEGER `lb:0,ub:69`
}
