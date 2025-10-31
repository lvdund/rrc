package ies

// RRC-InactiveConfig-r15 ::= SEQUENCE
type RrcInactiveconfigR15 struct {
	FulliRntiR15               IRntiR15
	ShortiRntiR15              ShortiRntiR15
	RanPagingcycleR15          *RrcInactiveconfigR15RanPagingcycleR15
	RanNotificationareainfoR15 *RanNotificationareainfoR15
	PeriodicRnauTimerR15       *RrcInactiveconfigR15PeriodicRnauTimerR15
	NexthopchainingcountR15    *Nexthopchainingcount
	Dummy                      *RrcInactiveconfigR15Dummy
}
