package ies

import "rrc/utils"

// UL-ConfigCommonList-NB-v1530 ::= SEQUENCE OF UL-ConfigCommon-NB-v1530
// SIZE (1.. maxNonAnchorCarriers-NB-r14)
type UlConfigcommonlistNbV1530 struct {
	Value utils.Sequence[UlConfigcommonNbV1530]
}
