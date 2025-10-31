package ies

// RAN-NotificationAreaInfo-r15 ::= CHOICE
const (
	RanNotificationareainfoR15ChoiceNothing = iota
	RanNotificationareainfoR15ChoiceCelllist
	RanNotificationareainfoR15ChoiceRanAreaconfiglist
)

type RanNotificationareainfoR15 struct {
	Choice            uint64
	Celllist          *PlmnRanAreacelllistR15
	RanAreaconfiglist *PlmnRanAreaconfiglistR15
}
