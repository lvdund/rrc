package ies

import "rrc/utils"

// SystemInformationBlockType13-r9 ::= SEQUENCE
// Extensible
type Systeminformationblocktype13R9 struct {
	MbsfnAreainfolistR9      MbsfnAreainfolistR9
	NotificationconfigR9     MbmsNotificationconfigR9
	Latenoncriticalextension *utils.OCTETSTRING
}
