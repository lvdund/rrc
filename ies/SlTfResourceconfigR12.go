package ies

import "rrc/utils"

// SL-TF-ResourceConfig-r12 ::= SEQUENCE
type SlTfResourceconfigR12 struct {
	PrbNumR12          utils.INTEGER `lb:0,ub:100`
	PrbStartR12        utils.INTEGER `lb:0,ub:99`
	PrbEndR12          utils.INTEGER `lb:0,ub:99`
	OffsetindicatorR12 SlOffsetindicatorR12
	SubframebitmapR12  SubframebitmapslR12
}
