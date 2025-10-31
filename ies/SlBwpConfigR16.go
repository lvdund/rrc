package ies

import "rrc/utils"

// SL-BWP-Config-r16 ::= SEQUENCE
// Extensible
type SlBwpConfigR16 struct {
	SlBwpId                BwpId
	SlBwpGenericR16        *SlBwpGenericR16
	SlBwpPoolconfigR16     *SlBwpPoolconfigR16
	SlBwpPoolconfigpsR17   *utils.Setuprelease[SlBwpPoolconfigR16]     `ext`
	SlBwpDiscpoolconfigR17 *utils.Setuprelease[SlBwpDiscpoolconfigR17] `ext`
}
