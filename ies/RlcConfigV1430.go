package ies

import "rrc/utils"

// RLC-Config-v1430 ::= CHOICE
type RlcConfigV1430 interface {
	isRlcConfigV1430()
}

type RlcConfigV1430Release struct {
	Value struct{}
}

func (*RlcConfigV1430Release) isRlcConfigV1430() {}

type RlcConfigV1430Setup struct {
	Value interface{}
}

func (*RlcConfigV1430Setup) isRlcConfigV1430() {}
