package ies

import "rrc/utils"

// DL-ConfigCommonList-NB-r14 ::= SEQUENCE OF DL-ConfigCommon-NB-r14
// SIZE (1.. maxNonAnchorCarriers-NB-r14)
type DlConfigcommonlistNbR14 struct {
	Value []DlConfigcommonNbR14
}
