package ies

import "rrc/utils"

// PUCCH-ConfigCommon-v1310 ::= SEQUENCE
type PucchConfigcommonV1310 struct {
	N1pucchAnInfolistR13              *N1pucchAnInfolistR13
	PucchNumrepetitionceMsg4Level0R13 *utils.ENUMERATED
	PucchNumrepetitionceMsg4Level1R13 *utils.ENUMERATED
	PucchNumrepetitionceMsg4Level2R13 *utils.ENUMERATED
	PucchNumrepetitionceMsg4Level3R13 *utils.ENUMERATED
}
