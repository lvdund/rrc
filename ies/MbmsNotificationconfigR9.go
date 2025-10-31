package ies

import "rrc/utils"

// MBMS-NotificationConfig-r9 ::= SEQUENCE
type MbmsNotificationconfigR9 struct {
	NotificationrepetitioncoeffR9 MbmsNotificationconfigR9NotificationrepetitioncoeffR9
	NotificationoffsetR9          utils.INTEGER `lb:0,ub:10`
	NotificationsfIndexR9         utils.INTEGER `lb:0,ub:6`
}
