package ies

import "rrc/utils"

// RAN-NotificationAreaInfo-r15 ::= CHOICE
type RanNotificationareainfoR15 interface {
	isRanNotificationareainfoR15()
}

type RanNotificationareainfoR15Celllist struct {
	Value PlmnRanAreacelllistR15
}

func (*RanNotificationareainfoR15Celllist) isRanNotificationareainfoR15() {}

type RanNotificationareainfoR15RanAreaconfiglist struct {
	Value PlmnRanAreaconfiglistR15
}

func (*RanNotificationareainfoR15RanAreaconfiglist) isRanNotificationareainfoR15() {}
