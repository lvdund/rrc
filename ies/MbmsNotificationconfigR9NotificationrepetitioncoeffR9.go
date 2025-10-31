package ies

import "rrc/utils"

// MBMS-NotificationConfig-r9-notificationRepetitionCoeff-r9 ::= ENUMERATED
type MbmsNotificationconfigR9NotificationrepetitioncoeffR9 struct {
	Value utils.ENUMERATED
}

const (
	MbmsNotificationconfigR9NotificationrepetitioncoeffR9EnumeratedNothing = iota
	MbmsNotificationconfigR9NotificationrepetitioncoeffR9EnumeratedN2
	MbmsNotificationconfigR9NotificationrepetitioncoeffR9EnumeratedN4
)
