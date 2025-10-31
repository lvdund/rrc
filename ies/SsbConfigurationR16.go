package ies

import "rrc/utils"

// SSB-Configuration-r16 ::= SEQUENCE
type SsbConfigurationR16 struct {
	SsbFreqR16              ArfcnValuenr
	HalfframeindexR16       SsbConfigurationR16HalfframeindexR16
	SsbsubcarrierspacingR16 Subcarrierspacing
	SsbPeriodicityR16       *SsbConfigurationR16SsbPeriodicityR16
	Sfn0OffsetR16           *SsbConfigurationR16Sfn0OffsetR16
	SfnSsbOffsetR16         utils.INTEGER  `lb:0,ub:15`
	SsPbchBlockpowerR16     *utils.INTEGER `lb:0,ub:50`
}
