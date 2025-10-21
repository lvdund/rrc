package ies

import "rrc/utils"

// DMRS-Config-v1310 ::= SEQUENCE
type DmrsConfigV1310 struct {
	DmrsTablealtR13 *utils.ENUMERATED
}
