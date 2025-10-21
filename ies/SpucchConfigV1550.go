package ies

import "rrc/utils"

// SPUCCH-Config-v1550 ::= CHOICE
type SpucchConfigV1550 interface {
	isSpucchConfigV1550()
}

type SpucchConfigV1550Release struct {
	Value struct{}
}

func (*SpucchConfigV1550Release) isSpucchConfigV1550() {}

type SpucchConfigV1550Setup struct {
	Value interface{}
}

func (*SpucchConfigV1550Setup) isSpucchConfigV1550() {}
