package ies

import "rrc/utils"

// PDCCH-ConfigCommon-followUnifiedTCI-State-v1720 ::= ENUMERATED
type PdcchConfigcommonFollowunifiedtciStateV1720 struct {
	Value utils.ENUMERATED
}

const (
	PdcchConfigcommonFollowunifiedtciStateV1720EnumeratedNothing = iota
	PdcchConfigcommonFollowunifiedtciStateV1720EnumeratedEnabled
)
