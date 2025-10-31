package ies

// RAN-NotificationAreaInfo ::= CHOICE
// Extensible
const (
	RanNotificationareainfoChoiceNothing = iota
	RanNotificationareainfoChoiceCelllist
	RanNotificationareainfoChoiceRanAreaconfiglist
)

type RanNotificationareainfo struct {
	Choice            uint64
	Celllist          *PlmnRanAreacelllist
	RanAreaconfiglist *PlmnRanAreaconfiglist
}
