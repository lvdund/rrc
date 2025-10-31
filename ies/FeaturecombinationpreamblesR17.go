package ies

import "rrc/utils"

// FeatureCombinationPreambles-r17 ::= SEQUENCE
// Extensible
type FeaturecombinationpreamblesR17 struct {
	FeaturecombinationR17                      FeaturecombinationR17
	StartpreambleforthispartitionR17           utils.INTEGER  `lb:0,ub:63`
	NumberofpreamblesperssbForthispartitionR17 utils.INTEGER  `lb:0,ub:64`
	SsbSharedroMaskindexR17                    *utils.INTEGER `lb:0,ub:15`
	GroupbconfiguredR17                        *FeaturecombinationpreamblesR17GroupbconfiguredR17
	SeparatemsgaPuschConfigR17                 *MsgaPuschConfigR16
	MsgaRsrpThresholdR17                       *RsrpRange
	RsrpThresholdssbR17                        *RsrpRange
	DeltapreambleR17                           *utils.INTEGER `lb:0,ub:6`
}
