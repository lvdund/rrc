package ies

import "rrc/utils"

// RLC-Config-v1530 ::= CHOICE
type RlcConfigV1530 interface {
	isRlcConfigV1530()
}

type RlcConfigV1530Release struct {
	Value struct{}
}

func (*RlcConfigV1530Release) isRlcConfigV1530() {}

type RlcConfigV1530Setup struct {
	Value interface{}
}

func (*RlcConfigV1530Setup) isRlcConfigV1530() {}
