package ies

import "rrc/utils"

// FeatureCombination-r17 ::= SEQUENCE
type FeaturecombinationR17 struct {
	RedcapR17          *utils.BOOLEAN
	SmalldataR17       *utils.BOOLEAN
	NsagR17            *NsagListR17
	Msg3RepetitionsR17 *utils.BOOLEAN
	Spare4             *utils.BOOLEAN
	Spare3             *utils.BOOLEAN
	Spare2             *utils.BOOLEAN
	Spare1             *utils.BOOLEAN
}
