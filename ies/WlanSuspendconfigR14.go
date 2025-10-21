package ies

import "rrc/utils"

// WLAN-SuspendConfig-r14 ::= SEQUENCE
type WlanSuspendconfigR14 struct {
	WlanSuspendresumeallowedR14        *bool
	WlanSuspendtriggersstatusreportR14 *bool
}
