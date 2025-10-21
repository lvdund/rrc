package ies

import "rrc/utils"

// UAC-Param-NB-r16 ::= CHOICE
type UacParamNbR16 interface {
	isUacParamNbR16()
}

type UacParamNbR16UacBarringcommon struct {
	Value UacBarringNbR16
}

func (*UacParamNbR16UacBarringcommon) isUacParamNbR16() {}

type UacParamNbR16UacBarringperplmnList struct {
	Value interface{}
}

func (*UacParamNbR16UacBarringperplmnList) isUacParamNbR16() {}
