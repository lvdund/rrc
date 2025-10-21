package ies

import "rrc/utils"

// RLC-Config ::= CHOICE
// Extensible
type RlcConfig interface {
	isRlcConfig()
}

type RlcConfigAm struct {
	Value interface{}
}

func (*RlcConfigAm) isRlcConfig() {}

type RlcConfigUmBiDirectional struct {
	Value interface{}
}

func (*RlcConfigUmBiDirectional) isRlcConfig() {}

type RlcConfigUmUniDirectionalUl struct {
	Value interface{}
}

func (*RlcConfigUmUniDirectionalUl) isRlcConfig() {}

type RlcConfigUmUniDirectionalDl struct {
	Value interface{}
}

func (*RlcConfigUmUniDirectionalDl) isRlcConfig() {}
