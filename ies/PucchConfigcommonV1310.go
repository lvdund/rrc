package ies

// PUCCH-ConfigCommon-v1310 ::= SEQUENCE
type PucchConfigcommonV1310 struct {
	N1pucchAnInfolistR13              *N1pucchAnInfolistR13
	PucchNumrepetitionceMsg4Level0R13 *PucchConfigcommonV1310PucchNumrepetitionceMsg4Level0R13
	PucchNumrepetitionceMsg4Level1R13 *PucchConfigcommonV1310PucchNumrepetitionceMsg4Level1R13
	PucchNumrepetitionceMsg4Level2R13 *PucchConfigcommonV1310PucchNumrepetitionceMsg4Level2R13
	PucchNumrepetitionceMsg4Level3R13 *PucchConfigcommonV1310PucchNumrepetitionceMsg4Level3R13
}
