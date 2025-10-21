package ies

import "rrc/utils"

// RSRQ-RangeConfig-r12 ::= CHOICE
type RsrqRangeconfigR12 interface {
	isRsrqRangeconfigR12()
}

type RsrqRangeconfigR12Release struct {
	Value struct{}
}

func (*RsrqRangeconfigR12Release) isRsrqRangeconfigR12() {}

type RsrqRangeconfigR12Setup struct {
	Value RsrqRangeV1250
}

func (*RsrqRangeconfigR12Setup) isRsrqRangeconfigR12() {}
