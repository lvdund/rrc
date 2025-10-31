package ies

import "rrc/utils"

// FeatureSetDownlink-v1700-mTRP-PDCCH-Repetition-r17 ::= SEQUENCE
type FeaturesetdownlinkV1700MtrpPdcchRepetitionR17 struct {
	NumbdTwopdcchR17  utils.INTEGER `lb:0,ub:3`
	MaxnumoverlapsR17 FeaturesetdownlinkV1700MtrpPdcchRepetitionR17MaxnumoverlapsR17
}
