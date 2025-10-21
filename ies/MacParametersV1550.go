package ies

import "rrc/utils"

// MAC-Parameters-v1550 ::= SEQUENCE
type MacParametersV1550 struct {
	ElcidSupportR15 *utils.ENUMERATED
}
