package ies

import "rrc/utils"

// NotificationMessageSidelink-r17-IEs ::= SEQUENCE
type NotificationmessagesidelinkR17 struct {
	IndicationtypeR17        *NotificationmessagesidelinkR17IesIndicationtypeR17
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *NotificationmessagesidelinkR17IesNoncriticalextension
}
