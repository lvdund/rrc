package ies

import "rrc/utils"

// PUCCH-ConfigCommon-v1430-pucch-NumRepetitionCE-Msg4-Level3-r14 ::= ENUMERATED
type PucchConfigcommonV1430PucchNumrepetitionceMsg4Level3R14 struct {
	Value utils.ENUMERATED
}

const (
	PucchConfigcommonV1430PucchNumrepetitionceMsg4Level3R14EnumeratedNothing = iota
	PucchConfigcommonV1430PucchNumrepetitionceMsg4Level3R14EnumeratedN64
	PucchConfigcommonV1430PucchNumrepetitionceMsg4Level3R14EnumeratedN128
)
