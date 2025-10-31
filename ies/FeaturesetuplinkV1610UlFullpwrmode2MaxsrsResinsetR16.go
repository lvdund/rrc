package ies

import "rrc/utils"

// FeatureSetUplink-v1610-ul-FullPwrMode2-MaxSRS-ResInSet-r16 ::= ENUMERATED
type FeaturesetuplinkV1610UlFullpwrmode2MaxsrsResinsetR16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610UlFullpwrmode2MaxsrsResinsetR16EnumeratedNothing = iota
	FeaturesetuplinkV1610UlFullpwrmode2MaxsrsResinsetR16EnumeratedN1
	FeaturesetuplinkV1610UlFullpwrmode2MaxsrsResinsetR16EnumeratedN2
	FeaturesetuplinkV1610UlFullpwrmode2MaxsrsResinsetR16EnumeratedN4
)
