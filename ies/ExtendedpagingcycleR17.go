package ies

import "rrc/utils"

// ExtendedPagingCycle-r17 ::= ENUMERATED
type ExtendedpagingcycleR17 struct {
	Value utils.ENUMERATED
}

const (
	ExtendedpagingcycleR17EnumeratedNothing = iota
	ExtendedpagingcycleR17EnumeratedRf256
	ExtendedpagingcycleR17EnumeratedRf512
	ExtendedpagingcycleR17EnumeratedRf1024
	ExtendedpagingcycleR17EnumeratedSpare1
)
