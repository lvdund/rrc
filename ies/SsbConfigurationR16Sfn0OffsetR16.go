package ies

import "rrc/utils"

// SSB-Configuration-r16-sfn0-Offset-r16 ::= SEQUENCE
type SsbConfigurationR16Sfn0OffsetR16 struct {
	SfnOffsetR16             utils.INTEGER  `lb:0,ub:1023`
	IntegersubframeoffsetR16 *utils.INTEGER `lb:0,ub:9`
}
