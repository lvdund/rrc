package ies

import "rrc/utils"

// MBMS-NotificationConfig-v1430 ::= SEQUENCE
type MbmsNotificationconfigV1430 struct {
	NotificationsfIndexV1430 utils.INTEGER `lb:0,ub:10`
}
