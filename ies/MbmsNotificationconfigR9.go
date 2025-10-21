package ies

import "rrc/utils"

// MBMS-NotificationConfig-r9 ::= SEQUENCE
type MbmsNotificationconfigR9 struct {
	NotificationrepetitioncoeffR9 utils.ENUMERATED
	NotificationoffsetR9          utils.INTEGER
	NotificationsfIndexR9         utils.INTEGER
}
