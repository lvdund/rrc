package ies

import "rrc/utils"

// FeatureSetUplink-v1710-mTRP-PUCCH-IntraSlot-r17 ::= ENUMERATED
type FeaturesetuplinkV1710MtrpPucchIntraslotR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1710MtrpPucchIntraslotR17EnumeratedNothing = iota
	FeaturesetuplinkV1710MtrpPucchIntraslotR17EnumeratedPf0_2
	FeaturesetuplinkV1710MtrpPucchIntraslotR17EnumeratedPf1_3_4
	FeaturesetuplinkV1710MtrpPucchIntraslotR17EnumeratedPf0_4
)
