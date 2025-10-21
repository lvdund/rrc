package ies

import "rrc/utils"

// RRC-InactiveConfig-r15 ::= SEQUENCE
type RrcInactiveconfigR15 struct {
	FulliRntiR15               IRntiR15
	ShortiRntiR15              ShortiRntiR15
	RanPagingcycleR15          *utils.ENUMERATED
	RanNotificationareainfoR15 *RanNotificationareainfoR15
	PeriodicRnauTimerR15       *utils.ENUMERATED
	NexthopchainingcountR15    *Nexthopchainingcount
	Dummy                      *interface{}
}
